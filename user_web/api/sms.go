package api

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strings"
	"study_mxshop_api/user_web/forms"
	"study_mxshop_api/user_web/global"
	"time"
)

func GenerateSmsCode(witdh int) string {
	// 生成width长度的短信验证码

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	if sendSmsForm.Type != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "短信类型不正确",
		})
		return
	}
	smsCode := GenerateSmsCode(6)

	if false { // 学习过程中不走下面逻辑
		client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.ServerConfig.AliSmsInfo.ApiKey, global.ServerConfig.AliSmsInfo.ApiSecrect)
		if err != nil {
			panic(err)
		}
		request := requests.NewCommonRequest()
		request.Method = "POST"
		request.Scheme = "https"
		request.Domain = "dysmsapi.aliyuncs.com"
		request.Version = "2017-05-25"
		request.ApiName = "SendSms"
		request.QueryParams["RegionId"] = "cn-beijing"
		request.QueryParams["PhoneNumbers"] = sendSmsForm.Mobile            // 手机号
		request.QueryParams["SignName"] = "慕学在线"                            // 阿里云验证过的项目名 自己设置
		request.QueryParams["TemplateCode"] = "SMS_181850725"               // 阿里云的短信模板号 自己设置
		request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" // 短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
		response, err := client.ProcessCommonRequest(request)
		fmt.Print(client.DoAction(request, response))
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	// 将验证码保存起来，用于手机号登录 - redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
		Password: global.ServerConfig.RedisInfo.Password,
	})
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)
	zap.S().Debugf("短信验证码: %s", smsCode)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
