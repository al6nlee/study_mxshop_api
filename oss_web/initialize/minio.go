package initialize

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"study_mxshop_api/oss_web/global"
)

func InitMinio() {
	ossInfo := global.ServerConfig.OssInfo
	endpoint := fmt.Sprintf("%s:%d", ossInfo.Host, ossInfo.PORT)
	accessKeyID := ossInfo.ApiKey
	secretAccessKey := ossInfo.ApiSecrect
	useSSL := false          // 使用 http 连接
	location := "cn-north-1" // 时区，cn-north-1 表示北京时间，cn-northwest-1 表示宁夏时间

	// 初始化 Minio 客户端
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		zap.S().Errorf("创建 MinIO 客户端失败: %v", err)
		return
	}

	// 设置全局 MinIO 客户端
	global.MinioClient = client

	// 获取桶名并尝试创建
	bucketName := ossInfo.BucketName
	ctx := context.Background()

	// 尝试创建桶
	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 如果桶已经存在，跳过创建操作
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists != nil {
			zap.S().Errorf("检查桶是否存在时发生错误: %v", errBucketExists)
			return
		}
		if exists {
			zap.S().Infof("桶 %s 已经存在，跳过创建", bucketName)
			return
		}
		// 其他错误，直接日志输出并终止
		zap.S().Errorf("创建桶 %s 时发生错误: %v", bucketName, err)
		return
	}
	zap.S().Infof("成功创建桶: %s", bucketName)
}
