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

package proff

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

//
// profile
//
type profile struct {
	from          *string
	replyTo       []*string
	returnPath    *string
	returnPathArn *string
	sourceArn     *string
}

//
// email struct
//
type Email struct {
	ses      *ses.SES
	profiles map[string]*profile
}

//
// Setup a profile to use with Send
//
func (this *Email) SetSetupProfile(name string, from string, replyTo []string, returnPath string, returnPathArn string, sourceArn string) bool {

	//
	// map
	//
	this.profiles = map[string]*profile{}

	//
	// profiles
	//
	this.profiles[name] = &profile{

		from:          aws.String(from),
		replyTo:       []*string{},
		returnPath:    aws.String(returnPath),
		returnPathArn: aws.String(returnPathArn),
		sourceArn:     aws.String(sourceArn),
	}

	//
	// for
	//
	for _, d := range replyTo {

		this.profiles[name].replyTo = append(this.profiles[name].replyTo, aws.String(d))
	}

	return true
}

//
//
//
func SetProfile(from string, info string) *profile {

	From := info + " <" + from + ">"
	ReturnPathx := "arn:aws:ses:us-east-1:873761630739:identity/" + from
	ReturnPathxArm := "arn:aws:ses:us-east-1:873761630739:identity/" + from

	//
	// config email
	//
	sender := new(Email)

	//fmt.Println(sender)

	sender.SetSetupProfile("default", From, []string{from},
		from,
		ReturnPathx,
		ReturnPathxArm)

	pr := sender.profiles["default"]

	fmt.Println("From: ", pr.from)

	if pr == nil {

		fmt.Println("Error profiles: ", pr)
		os.Exit(1)
	}

	return pr
}
