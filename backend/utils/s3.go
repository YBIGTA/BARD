package util

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3 *S3Info

type S3Info struct {
	BucketName string
	Region     string
	Client     *s3.Client
}

func InitializeS3Client() {

	creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	S3 = &S3Info{
		BucketName: os.Getenv("AWS_BUCKET_NAME"),
		Region:     os.Getenv("AWS_REGION"),
		Client:     client,
	}

}

func (s *S3Info) UploadFile(file multipart.File, object_name string) error {

	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(object_name),
		Body:   file,
	})

	return err
}

func (s *S3Info) GetFileUrl(object_name string) string {

	return "https://" + s.BucketName + ".s3." + s.Region + ".amazonaws.com/" + object_name
}
