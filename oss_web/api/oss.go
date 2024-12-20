package api

import (
	"github.com/gin-gonic/gin"
	"study_mxshop_api/oss_web/global"
	"time"
)

func PresignedPut(c *gin.Context) {
	// 1. 优先从请求中获取 bucketName，如果没有提供则生成一个唯一的 bucketName
	bucketName := c.DefaultQuery("bucketName", "")
	if bucketName == "" {
		// 如果没有提供，则生成一个唯一的 bucketName
		bucketName = "bucket_" + time.Now().Format("20060102150405")
	}
	bucketName = global.ServerConfig.OssInfo.Prefix + "/" + bucketName

	presignedURL, err := global.MinioClient.PresignedPutObject(c,
		global.ServerConfig.OssInfo.BucketName,
		bucketName,
		time.Minute*15,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "生成预签名失败", "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"presigned_url": presignedURL.String(),
	})
}

// 处理 MinIO 上传事件回调
func HandlerRequest(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok!"})
}
