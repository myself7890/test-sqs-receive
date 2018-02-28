package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIWXVSIDZT7YVCRAA")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "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")

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	client := sqs.New(sess)

	queueUrl := "https://sqs.us-east-1.amazonaws.com/127686158421/Test.fifo"
	ListenToQueue(client, queueUrl)
}

func ListenToQueue(client *sqs.SQS, url string) error {
	for {
		log.Println("Try to pull message...")
		result, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(url),
			MaxNumberOfMessages: aws.Int64(10),
			MessageAttributeNames: aws.StringSlice([]string{
				"All",
			}),
			WaitTimeSeconds: aws.Int64(5),
		})

		if err != nil {
			log.Println(err.Error())
			return err
		}

		log.Println("YAAAA", result.Messages)
		for _, message := range result.Messages {
			_, err := client.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      &url,
				ReceiptHandle: message.ReceiptHandle,
			})

			if err != nil {
				log.Println("Delete error", err)
			}
		}
	}
	return nil
}
