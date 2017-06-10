# goses

Sending emails with aws ses, a package to make it easier to understand and use aws sdk to send emails.

We did something simple that we are still improving, for every email we have to call the method.

Emails with copy and hidden will be implemented in another method.

We created a package called gses (golang SES), which creates and mounts all objects so that we can send our email successfully.

In order for your email to be sent successfully using the aws sdk you need to have an email validated by SES (Verify This Email Address), console access can be done by clicking here https://console.aws.amazon.com/ses, 
it will Also need your Identity ARN.

# Verify This Email Address - SES Example

![image](https://github.com/jeffotoni/goses/blob/master/img/identity-arn.png)


# Examples of submissions

```go

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
// EmailCC 		:= "emailCc@domain.com"
// EmailBc 		:= "emailBcc@domain.com"
// Html 		:= "<html><body><h1>test html context</h1></body></html>"
// Subject 		:= "Your message title"
//
err = S.Send(
		"emailTo@domain.com", 

		"emailCc@domain.com", 

		"emailBcc@domain.com", 

		"<h1>Test send email....</h1>", 

		"Test send email to me goses 1002",

		)
if err != nil {

	fmt.Printf("Error %v\n", err)

} else {

	fmt.Println("Send Sucess")
}

//
// EmailTo 		:= "emailTo1@domain.com,emailTo2@domain.com,emailTo3@domain.com"
// EmailCC 		:= "emailCc1@domain.com,emailCc2@domain.com"
// EmailBc 		:= "emailBcc1@domain.com,emailBcc2@domain.com"
// Html 		:= "<html><body><h1>test html context</h1></body></html>"
// Subject 		:= "Your message title"
//
err = S.Send(
	// Multiple To emails separated by commas
	" emailTo@domain.com, email2@domain.com, email3@domain.com ",

	// Multiple Cc emails separated by commas
	"emailCc1@domain.com,emailCc2@domain.com", 

	// Multiple Bcc emails separated by commas
	"emailBcc1@domain.com,emailBcc2@domain.com", 

	// Html Text
	"<h1>Test send email....</h1>",

	// Subject here
	"Test send email to me goses 1002",

) if err != nil {

	fmt.Printf("Error %v\n", err)

} else {

	fmt.Println("Send Sucess")
}

```
