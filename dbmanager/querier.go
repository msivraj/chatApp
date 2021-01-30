package dbmanager

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateNewDbSession() {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "msivraj"),
		HTTPClient: NewHTTPClientWithSettings(HTTPClientSettings{
			Connect:          5 * time.Second,
			ExpectContinue:   1 * time.Second,
			IdleConn:         90 * time.Second,
			ConnKeepAlive:    30 * time.Second,
			MaxAllIdleConns:  100,
			MaxHostIdleConns: 10,
			ResponseHeader:   5 * time.Second,
			TLSHandshake:     5 * time.Second,
		}),
	})

	svc := dynamodb.New(sess)

	input := &dynamodb.ListTablesInput{}

	fmt.Print("Tables: \n")

	for {
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		for _, n := range result.TableNames {
			fmt.Println(*n)
		}

		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}

	}

	// //every service that is made from a session has its handlers on it
	// // sess.Handlers.Send.PushFront(func(r *request.Request) {
	// // 	// Log every request made and its payload
	// // 	fmt.Print("Request: ", r.ClientInfo.ServiceName, r.Operation, " Payload: ", r.Params)
	// // })

	// //get service from session, session should be a global one value unless its config needs to change
	// // svc := s3.New(sess)

	// // use when just want to use default values
	// // sess := session.Must(session.NewSessionWithOptions(session.Options{
	// // 	SharedConfigState: session.SharedConfigEnable,
	// // }))

	// //copying a session region is overridden by what is passed in
	// // usEast2Sess := sess.Copy(&aws.Config{Region: aws.String("us-east-2")})

	// fmt.Print(err, "\n")
	// fmt.Print(sess, "\n")

	// svc := s3.New(sess)

}
