package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
)

// func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	_, err = uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(fileName),
// 		Body:   file,
// 	})

// 	return err
// }

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-west-2"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	bucketName := "firstbucket2345678980"
	uploader := s3manager.NewUploader(sess)
	filename := "data.txt"

	err = UploadFile(uploader, "data.txt", bucketName, filename)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}

	fmt.Println("Successfully uploaded file!")
}
