package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	return err
}

func main() {

	fmt.Print("Enter your Data: \n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_str := scanner.Text()
	//fmt.Println("Input value is: ", input_str)

	myfile, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(myfile)

	_, err2 := myfile.WriteString(input_str)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Done Writing to file")

	content, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	datain := fmt.Sprintf("%s", content)

	fmt.Println("datain is:", datain)

	defer myfile.Close()

	//
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
