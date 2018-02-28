package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	os.Setenv("AWS_ACCESS_KEY_ID", "INSERT ACCESS KEY IS HERE")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "INSERT ACCESS KEY HERE")

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	client := sqs.New(sess)

	queueUrl := "https://sqs.us-east-1.amazonaws.com/127686158421/Test.fifo"

	result, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: aws.Int64(5),
	})

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("YAAAA", result.Messages)
}
