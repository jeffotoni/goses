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

	gses "github.com/jeffotoni/goses/pkg"
	//gses "./pkg"
)

func main() {

	//
	//
	//
	var err error

	//
	// Region, IdentityArn, From, Info <email from>
	//
	// "us-east-1", "IdentityArn Here", "emailFrom@domain.com", "Lets test send email ses.."
	//
	S := gses.SetProfile(

		"us-east-1",

		"IdentityArn Here",

		"from@domain.com",

		"Info Lets test send email ses..",
	)

	//
	// EmailTo 		:= "emailTo@domain.com"
	// Html 		:= --
	// Subject 		:= --
	// EmailCC 		:= --
	// EmailBc 		:= --
	//

	err = S.Send(

		"emailTo@domain.com",
	)
	if err != nil {

		fmt.Printf("Error %v\n", err)

	} else {

		fmt.Println("Send Sucess")
	}

	//
	// EmailTo 		:= "emailTo@domain.com"
	// Html 		:= "<html><body><h1>test html context</h1></body></html>"
	// Subject 		:= "Your message title"
	// EmailCC 		:= --
	// EmailBc 		:= --
	//

	err = S.Send(

		"emailTo@domain.com",

		"<h1>Test send email....</h1>",

		"Test send email to me goses 1002",
	)
	if err != nil {

		fmt.Printf("Error %v\n", err)

	} else {

		fmt.Println("Send Sucess")
	}

	//
	// EmailTo 		:= "emailTo@domain.com"
	// Html 		:= "<html><body><h1>test html context</h1></body></html>"
	// Subject 		:= "Your message title"
	//
	// EmailCC 		:= "emailCc@domain.com"
	//

	err = S.Send(

		"emailTo@domain.com",

		"<h1>Test send email....</h1>",

		"Test send email to me goses 1002",

		"emailCc@domain.com",
	)
	if err != nil {

		fmt.Printf("Error %v\n", err)

	} else {

		fmt.Println("Send Sucess")
	}

	//
	// EmailTo 		:= "emailTo@domain.com"
	// Html 		:= "<html><body><h1>test html context</h1></body></html>"
	// Subject 		:= "Your message title"
	// EmailCC		:= ""
	// EmailBc 		:= "emailBcc@domain.com"
	//

	err = S.Send(

		"emailTo@domain.com",

		"<h1>Test send email....</h1>",

		"Test send email to me goses 1002",

		"",

		"emailBcc@domain.com",
	)
	if err != nil {

		fmt.Printf("Error %v\n", err)

	} else {

		fmt.Println("Send Sucess")
	}

	//
	// EmailTo 		:= "emailTo1@domain.com,emailTo2@domain.com,emailTo3@domain.com"
	// Html 		:= "<html><body><h1>test html context</h1></body></html>"
	// Subject 		:= "Your message title"
	//
	// EmailCC 		:= "emailCc1@domain.com,emailCc2@domain.com"
	// EmailBc 		:= "emailBcc1@domain.com,emailBcc2@domain.com"
	//
	err = S.Send(

		// Multiple To emails separated by commas
		" emailTo@domain.com, email2@domain.com, email3@domain.com ",

		// Html Text
		"<h1>Test send email....</h1>",

		// Subject here
		"Test send email to me goses 1002",

		// Multiple Cc emails separated by commas
		"emailCc1@domain.com,emailCc2@domain.com",

		// Multiple Bcc emails separated by commas
		"emailBcc1@domain.com,emailBcc2@domain.com",
	)

	if err != nil {

		fmt.Printf("Error %v\n", err)

	} else {

		fmt.Println("Send Sucess")
	}
}
