# goses

Sending emails with aws ses, a package to make it easier to understand and use aws sdk to send emails.

We did something simple that we are still improving, for every email we have to call the method.

Emails with copy and hidden will be implemented in another method.

We created a package called gses (golang SES), which creates and mounts all objects so that we can send our email successfully.

In order for your email to be sent successfully using the aws sdk you need to have an email validated by SES (Verify This Email Address), console access can be done by clicking here https://console.aws.amazon.com/ses, it will Also need your Identity RNA.

# Methods

```go

//
// Region , IdentityArn , From, Info <email from>
//
// Ie := &proff.InfEmail{"us-east-1", "xxxxxxxx", "email@domain.com", "Lets test send email ses.."}
//
S := gses.SetProfile("us-east-1", "xxxxxxxxx", "email@domain.com", "Lets test send email ses..")

//
// EmailTo 		:= "emailTo@domain.com"
// Html 		:= "<html><body><h1>test html context</h1></body></html>"
// Subject 		:= "Your message title"
//
err := S.Send("emailTo@domain.com", "<h1>Test send email....</h1>", "Test send email to me goses")
if err != nil {

	fmt.Printf("Error %s => %v\n", err)
}

```