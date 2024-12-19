package api

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"time"
)

// 初始化 MinIO 客户端
func NewMinioClient() *minio.Client {
	endpoint := "172.26.25.139:30081"
	accessKeyID := "laOaPuoMCOkUUE4eTHky"
	secretAccessKey := "W93aHwb1pZAlMmpD8lI0AzEmi3SBiHkHhsV8hlez"
	useSSL := false // http 连接

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal("Unable to create MinIO client:", err)
	}
	return client
}

func Token(c *gin.Context) {
	client := NewMinioClient()

	// 假设我们需要生成一个上传文件的预签名 URL
	bucketName := "oss-bucket"
	objectName := "20241220"
	expiry := time.Minute * 15 // 15分钟有效期

	// 生成上传的预签名 URL
	presignedURL, err := client.PresignedPutObject(c, bucketName, objectName, expiry)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to generate presigned URL", "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"presigned_url": presignedURL,
	})
}

// 定义事件记录结构体，匹配 MinIO 的事件通知格式
type NotificationRecord struct {
	EventVersion string `json:"eventVersion"`
	EventSource  string `json:"eventSource"`
	AwsRegion    string `json:"awsRegion"`
	EventTime    string `json:"eventTime"`
	EventName    string `json:"eventName"`
	UserIdentity struct {
		PrincipalID string `json:"principalId"`
	} `json:"userIdentity"`
	RequestParameters struct {
		SourceIPAddress string `json:"sourceIPAddress"`
	} `json:"requestParameters"`
	ResponseElements struct {
		XAmzRequestID string `json:"x-amz-request-id"`
		XAmzID2       string `json:"x-amz-id-2"`
	} `json:"responseElements"`
	S3 struct {
		SChemaVersion   string `json:"s3SchemaVersion"`
		ConfigurationID string `json:"configurationId"`
		Bucket          struct {
			Name string `json:"name"`
			ARN  string `json:"arn"`
		} `json:"bucket"`
		Object struct {
			Key  string `json:"key"`
			Size int64  `json:"size"`
		} `json:"object"`
	} `json:"s3"`
}

// 定义接收事件的结构体
type NotificationEvent struct {
	Records []NotificationRecord `json:"Records"`
}

// 处理 MinIO 上传事件回调
func HandlerRequest(c *gin.Context) {
	var event NotificationEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": "Invalid event data"})
		return
	}

	// 处理事件
	for _, record := range event.Records {
		log.Printf("Event Name: %s, Bucket: %s, Object: %s, Size: %d\n", record.EventName, record.S3.Bucket.Name, record.S3.Object.Key, record.S3.Object.Size)
	}

	c.JSON(200, gin.H{
		"message": "Event processed successfully",
	})
}
