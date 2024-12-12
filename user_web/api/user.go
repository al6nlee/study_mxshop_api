package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"study_mxshop_api/user_web/forms"
	"study_mxshop_api/user_web/global"
	reponse "study_mxshop_api/user_web/global/response"
	"study_mxshop_api/user_web/middlewares"
	"study_mxshop_api/user_web/models"
	"study_mxshop_api/user_web/proto"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func GetUserList(ctx *gin.Context) {
	// 通过jwt中间组件验证后，会set一些参数
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	// 获取Query传的参数
	pn := ctx.DefaultQuery("pn", "1")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	// 拨号连接srv
	ip := global.ServerConfig.UserSrv.Host
	port := global.ServerConfig.UserSrv.Port
	userConn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接srv失败", "err", err)
		ctx.JSON(200, gin.H{
			"msg": "连接srv失败",
		})
		return
	}
	// 生成grpc的client调用接口
	userSrvClient := proto.NewUserClient(userConn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("获取用户列表失败", "err", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := gin.H{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := reponse.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			// Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			Birthday: reponse.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func PassWordLogin(ctx *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil { // ShouldBind自动判断请求是form还是json并做绑定
		HandleValidatorError(ctx, err)
		return
	}

	// 效验账户密码前，先对验证码做验证
	if store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	// 登录的逻辑
	// 拨号连接srv
	ip := global.ServerConfig.UserSrv.Host
	port := global.ServerConfig.UserSrv.Port
	userConn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接srv失败", "err", err)
		ctx.JSON(200, gin.H{
			"msg": "连接srv失败",
		})
		return
	}
	// 生成grpc的client调用接口
	userSrvClient := proto.NewUserClient(userConn)
	rsp, err := userSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "用户不存在",
				})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
				return
			}
		}
	}
	// 只是查询到用户了而已，并没有检查密码
	passRsp, pasErr := userSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
		Password:          passwordLoginForm.PassWord,
		EncryptedPassword: rsp.PassWord,
	})
	if pasErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败",
		})
		return
	}
	if passRsp.Success {
		// 登录成功
		j := middlewares.NewJWT()
		claims := models.CustomClaims{
			ID:          uint(rsp.Id),
			NickName:    rsp.NickName,
			AuthorityId: uint(rsp.Role),
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),               // 签名的生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
				Issuer:    "SAAS",                          // 签发机构
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "生成token失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"id":         rsp.Id,
			"nick_name":  rsp.NickName,
			"token":      token,
			"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "密码错误",
		})
	}
}
