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

package gses

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

//
// email struct
// We will assemble our data map with this structure
//
type SesEmail struct {
	ses      *ses.SES
	Profiles map[string]*profile
}

//
// profile
//
type profile struct {
	From          *string
	Sfrom         string
	ReplyTo       []*string
	ReturnPath    *string
	ReturnPathArn *string
	SourceArn     *string
	Region        string
}

// func (sf *Inf) InfEmail() {

// }

//
// Setup a profile to use with Send
// With this function, we were able to make the subtitle work correctly
//
func (this *SesEmail) SetSetupProfile(name string, from string, replyTo []string, returnPath string, returnPathArn string, sourceArn string, region string) bool {

	//
	// or data map
	//
	this.Profiles = map[string]*profile{}

	//
	// profiles
	//
	this.Profiles[name] = &profile{

		From:          aws.String(from),
		Sfrom:         from,
		ReplyTo:       []*string{},
		ReturnPath:    aws.String(returnPath),
		ReturnPathArn: aws.String(returnPathArn),
		SourceArn:     aws.String(sourceArn),
		Region:        region,
	}

	//
	// for
	//
	for _, d := range replyTo {

		this.Profiles[name].ReplyTo = append(this.Profiles[name].ReplyTo, aws.String(d))
	}

	return true
}

//
//
//
func SetProfile(

	//
	// region aws ex: us-east-1
	//
	Region string,

	//
	// https://console.aws.amazon.com/ses
	// Identity ARN: arn:aws:ses:region-aws:xxxx:identity/yourmail@domain.com
	//
	IdentityArn string,

	//
	// Mail that it will send, it has to be configured on your SES
	//
	From string,

	//
	// Message that will be displayed in from ex: "text info here <emailfrom@domain.com>"
	//
	Info string) *profile {

	//
	// Identity ARN: arn:aws:ses:region-aws:xxxx:identity/yourmail@domain.com
	//
	IdentityARN := "arn:aws:ses:" + Region + ":" + IdentityArn + ":identity/" + From

	//
	//
	//
	FromX := Info + " <" + From + ">"

	//
	//
	//
	ReturnPathx := IdentityARN

	//
	//
	//
	ReturnPathxArm := IdentityARN

	//
	// config email
	//
	sender := new(SesEmail)

	//
	//
	//
	sender.SetSetupProfile("default", FromX, []string{From},
		From,
		ReturnPathx,
		ReturnPathxArm, Region)

	prof := sender.Profiles["default"]

	if prof == nil {

		fmt.Println("Error profiles: ", prof)
		os.Exit(1)
	}

	return prof
}

func (pf *profile) Send(EmailTo string, Cc string, Bc string, Html string, Subject string) error {

	DestinationV := &ses.Destination{}

	if Cc != "" && Bc == "" {

		vCc := []*string{

			aws.String(Cc), // Required
			// More values...
		}

		DestinationV = &ses.Destination{ // Required

			//BccAddresses: vBc,

			CcAddresses: vCc,

			ToAddresses: []*string{

				aws.String(EmailTo), // Required

				// More values...
			},
		}

	} else if Cc == "" && Bc != "" {

		vBc := []*string{

			aws.String(Bc), // Required
			// More values...
		}

		DestinationV = &ses.Destination{ // Required

			BccAddresses: vBc,

			//CcAddresses: vCc,

			ToAddresses: []*string{

				aws.String(EmailTo), // Required

				// More values...
			},
		}
	} else if Cc != "" && Bc != "" {

		vBc := []*string{

			aws.String(Bc), // Required
			// More values...
		}

		vCc := []*string{

			aws.String(Cc), // Required
			// More values...
		}

		DestinationV = &ses.Destination{ // Required

			BccAddresses: vBc,

			CcAddresses: vCc,

			ToAddresses: []*string{

				aws.String(EmailTo), // Required

				// More values...
			},
		}
	} else {

		DestinationV = &ses.Destination{ // Required

			ToAddresses: []*string{

				aws.String(EmailTo), // Required

				// More values...
			},
		}
	}

	params := &ses.SendEmailInput{

		Destination: DestinationV,

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

		Source:           pf.From,
		ReplyToAddresses: pf.ReplyTo,
		ReturnPath:       pf.ReturnPath,
		ReturnPathArn:    pf.ReturnPathArn,
		SourceArn:        pf.SourceArn,
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(pf.Region)},
	)

	svc := ses.New(sess)

	_, err = svc.SendEmail(params)

	if err != nil {

		//fmt.Printf("Error %s => %v\n", EmailTo, err)
		return err

	} else {

		//fmt.Printf("Send success %s\n", EmailTo)
		return nil
	}
}
