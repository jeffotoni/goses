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

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

package main

type profile struct {
	from          *string
	replyTo       []*string
	returnPath    *string
	returnPathArn *string
	sourceArn     *string
}

type Email struct {
	ses      *ses.SES
	profiles map[string]*profile
}

// Setup a profile to use with Send
func (this *Email) SetupProfile(name string, from string, replyTo []string, returnPath string, returnPathArn string, sourceArn string) bool {

	this.profiles = map[string]*profile{}

	this.profiles[name] = &profile{

		from:          aws.String(from),
		replyTo:       []*string{},
		returnPath:    aws.String(returnPath),
		returnPathArn: aws.String(returnPathArn),
		sourceArn:     aws.String(sourceArn),
	}

	for _, d := range replyTo {
		this.profiles[name].replyTo = append(this.profiles[name].replyTo, aws.String(d))
	}

	return true
}

func func main() {
	
	EmailTo := "jeff.otoni@s3wf.com.br"

	Html 	:= "<h1>Test send email....</h1>"

	Subject := "Test send email to me"

	params 	:= &ses.SendEmailInput{

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


	_, err := svc.SendEmail(params)

}
