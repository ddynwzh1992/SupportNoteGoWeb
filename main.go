// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// var router *gin.Engine

// func main() {

// 	// Set the router as the default one provided by Gin
// 	router = gin.Default()

// 	router.Static("/","./")
// 	// Start serving the application
// 	router.Run()
// }
package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.Static("/","./")

	router.POST("/login", func(c *gin.Context) {
		cc := c.PostForm("cc")
        subject := c.PostForm("subject")
        messages := c.PostForm("messages")
		//body,err:= c.GetRawData()

		body := fmt.Sprintf("The email is CC to %v, subject is %v, messages is %v",cc,subject,messages)
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("cn-northwest-1"),
			Credentials: credentials.NewStaticCredentials("AKIAPFNROODEA5EVB4YQ", "BIvLir98p8bP/+hlDmoi+gitCv47yQ5zjJMxfA9q", ""),
		})
		svc := sqs.New(sess)

		// URL to our queue
		qURL := "https://sqs.cn-northwest-1.amazonaws.com.cn/436307285021/supportnotetest"


		result, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"Title": &sqs.MessageAttributeValue{
					DataType:    aws.String("String"),
					StringValue: aws.String("The Whistler"),
				},
				"Author": &sqs.MessageAttributeValue{
					DataType:    aws.String("String"),
					StringValue: aws.String("John Grisham"),
				},
				"WeeksOn": &sqs.MessageAttributeValue{
					DataType:    aws.String("Number"),
					StringValue: aws.String("6"),
				},
			},
			MessageBody: aws.String(string(body)),
			QueueUrl:    &qURL,
		})

		fmt.Println("end")
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		fmt.Println("Success", *result.MessageId)
	})
	router.Run()
}

