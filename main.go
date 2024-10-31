package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	bucketName := os.Getenv("BUCKET_NAME")
	region := os.Getenv("AWS_REGION")
	endpoint := os.Getenv("AWS_ENDPOINT_URL_S3")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	domain := os.Getenv("DOMAIN")

	if bucketName == "" || region == "" || endpoint == "" || accessKeyID == "" || secretAccessKey == "" {
		log.Fatal("All required environment variables must be set: BUCKET_NAME, AWS_REGION, AWS_ENDPOINT_URL_S3, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY")
	}

	var fileName string
	var fileContent io.Reader

	if len(os.Args) > 1 {
		// Read from the provided filename
		fileName = os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()
		fileContent = file
	} else {
		// Read from standard input
		fileName = fmt.Sprintf("%d.txt", time.Now().Unix())
		fileContent = os.Stdin
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			if service == s3.ServiceID {
				return aws.Endpoint{
					URL: endpoint,
				}, nil
			}
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})),
	)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(s3Client)

	fileBuffer := new(bytes.Buffer)
	_, err = io.Copy(fileBuffer, fileContent)
	if err != nil {
		log.Fatalf("Failed to read file content: %v", err)
	}

	fileType := mime.TypeByExtension(filepath.Ext(fileName))
	if fileType == "" {
		fileType = "application/octet-stream"
	}

	uploadInput := &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &fileName,
		Body:        bytes.NewReader(fileBuffer.Bytes()),
		ContentType: &fileType,
		Expires:     aws.Time(time.Now().Add(10 * 24 * time.Hour)),
		ACL:         "public-read",
	}

	_, err = uploader.Upload(context.TODO(), uploadInput)
	if err != nil {
		log.Fatalf("Failed to upload file:%v", err)
	}

	url := fmt.Sprintf("%s/%s/%s", endpoint, bucketName, fileName)
	if domain != "" {
		url = fmt.Sprintf("https://%s/%s", domain, fileName)
	}

	err = clipboard.WriteAll(url)
	if err != nil {
		log.Fatalf("Failed to write URL to clipboard: %v", err)
	}

	fmt.Printf("%s\n", url)
}
