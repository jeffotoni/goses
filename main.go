/***********
*
*
* project server the Upload
*
* @package     main
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*
 */

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	prof "github.com/jeffotoni/goses/pkg/"
)

func main() {

	From := "xx@domain.com"
	Info := "Test send email for me"
	EmailTo := "your@mail.com"
	Html := "<h1>Test send email....</h1>"
	Subject := "Test send email to me"

	pr := prof.SetProfile(From, Info)

	params := &ses.SendEmailInput{

		Destination: &ses.Destination{ // Required
			// BccAddresses: []*string{
			//     aws.String("teste@s3wf.com.br"), // Required
			//     // More values...
			// },
			// CcAddresses: []*string{
			//     aws.String("teste@s3wf.com.br"), // Required
			//     // More values...
			// },
			ToAddresses: []*string{
				aws.String(EmailTo), // Required
				// More values...
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data:    aws.String(Html), // Required
					Charset: aws.String("utf-8"),
				},
				//,
				// Text: &ses.Content{
				//     Data:    aws.String("MessageData"), // Required
				//     Charset: aws.String("Charset"),
				// },
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(Subject), // Required
				Charset: aws.String("utf-8"),
			},
		},

		Source:           pr.from,
		ReplyToAddresses: pr.replyTo,
		ReturnPath:       pr.returnPath,
		ReturnPathArn:    pr.returnPathArn,
		SourceArn:        pr.sourceArn,

		//Source: aws.String(tmp_from),
		//, // Required
		// ReplyToAddresses: []*string{
		//     aws.String("Address"), // Required
		//     // More values...
		// },
		//ReturnPath:    aws.String("Address"),
		//ReturnPathArn: aws.String("AmazonResourceName"),
		//SourceArn:     aws.String("AmazonResourceName"),
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := ses.New(sess)

	_, err = svc.SendEmail(params)

	if err != nil {

		fmt.Println("Error %s => %v\n", EmailTo, err)

	} else {

		fmt.Println("Send success %s\n", EmailTo)
	}
}
