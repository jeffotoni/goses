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
	"strings"

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

//
// EmailTo string, Cc string, Bc string, Html string, Subject string
//
//func (pf *profile) Send(EmailTo string, Html string, Subject string, Cc string, Bcc string) error {
func (pf *profile) Send(paramses ...string) error {

	var EmailTo string
	var Cc string
	var Bcc string
	var Html string
	var Subject string

	if len(paramses) > 0 && len(paramses) <= 5 {

		if len(paramses) == 5 {

			EmailTo = paramses[0]

			Html = paramses[1]

			Subject = paramses[2]

			Cc = paramses[3]

			Bcc = paramses[4]

		} else if len(paramses) == 4 {

			EmailTo = paramses[0]

			Html = paramses[1]

			Subject = paramses[2]

			Cc = paramses[3]

			Bcc = ""

		} else if len(paramses) == 3 {

			EmailTo = paramses[0]

			Html = paramses[1]

			Subject = paramses[2]

			Cc = ""

			Bcc = ""

		} else if len(paramses) == 2 {

			EmailTo = paramses[0]

			Html = paramses[1]

			Subject = ""

			Cc = ""

			Bcc = ""

		} else if len(paramses) == 1 {

			EmailTo = paramses[0]

			Html = ""

			Subject = ""

			Cc = ""

			Bcc = ""
		}

	} else {

		fmt.Println("Error Parameters is missing")
		os.Exit(1)
	}

	//
	//
	//
	if EmailTo == "" {

		fmt.Println("Error EmailTo Required")
		os.Exit(1)
	}

	//
	//
	//
	DestinationV := &ses.Destination{}

	//
	//
	//
	ToAddressesMail := []*string{}

	//
	//
	//
	CcAddressesMail := []*string{}

	//
	//
	//
	BccAddressesMail := []*string{}

	//
	//
	//
	EmailTo = strings.Trim(EmailTo, " ")

	//
	//
	//
	arrayMailTo := strings.Split(EmailTo, ",")

	for i := range arrayMailTo {

		//
		//
		//
		mailClean := strings.TrimSpace(arrayMailTo[i])

		//
		//
		//
		ToAddressesMail = append(ToAddressesMail, aws.String(mailClean))
	}

	//
	//
	//
	if Cc != "" {

		//
		//
		//
		Cc = strings.Trim(Cc, " ")

		//
		//
		//
		arrayMailCc := strings.Split(Cc, ",")

		for i := range arrayMailCc {

			//
			//
			//
			mailCcClean := strings.TrimSpace(arrayMailCc[i])

			//
			//
			//
			CcAddressesMail = append(CcAddressesMail, aws.String(mailCcClean))
		}
	}

	//
	//
	//
	if Bcc != "" {

		//
		//
		//
		Bcc = strings.Trim(Bcc, " ")

		//
		//
		//
		arrayMailBCc := strings.Split(Bcc, ",")

		for i := range arrayMailBCc {

			//
			//
			//
			mailBccClean := strings.TrimSpace(arrayMailBCc[i])

			//
			//
			//
			BccAddressesMail = append(BccAddressesMail, aws.String(mailBccClean))
		}
	}

	//
	//
	//
	if Cc != "" && Bcc == "" {

		DestinationV = &ses.Destination{ // Required

			CcAddresses: CcAddressesMail,

			ToAddresses: ToAddressesMail,
		}

	} else if Cc == "" && Bcc != "" {

		DestinationV = &ses.Destination{ // Required

			BccAddresses: BccAddressesMail,

			ToAddresses: ToAddressesMail,
		}
	} else if Cc != "" && Bcc != "" {

		DestinationV = &ses.Destination{ // Required

			BccAddresses: BccAddressesMail,

			CcAddresses: CcAddressesMail,

			ToAddresses: ToAddressesMail,
		}
	} else {

		DestinationV = &ses.Destination{ // Required

			ToAddresses: ToAddressesMail,
		}
	}

	params := &ses.SendEmailInput{

		//
		//
		//
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
