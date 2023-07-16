package s3

import "github.com/aws/aws-sdk-go/service/s3/s3iface"

type S3Session interface {
	GetClient() s3iface.S3API
}
