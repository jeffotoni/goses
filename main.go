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
	import gses "github.com/jeffotoni/goses/pkg"
	//gses "./pkg"
)

func main() {

	//
	// Region , IdentityArn , From, Info <email from>
	//
	//Ie := &proff.InfEmail{"us-east-1", "xxxxxxxx", "email@domain.com", "Lets test send email ses.."}
	//
	S := gses.SetProfile("us-east-1", "xxxxxxxxx", "email@domain.com", "Lets test send email ses..")

	//
	// EmailTo 		:= "emailTo@domain.com"
	// Html 		:= "<html><body><h1>test html context</h1></body></html>"
	// Subject 		:= "Your message title"
	//
	err := S.Send("emailTo@domain.com", "<h1>Test send email....</h1>", "Test send email to me goses 1000")
	if err != nil {

		fmt.Printf("Error %s => %v\n", err)
	}
}
