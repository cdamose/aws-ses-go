package main

import (
	"log"
	"os"

	"github/cdamose/aws-ses-go/email"
	"github.com/joho/godotenv"
)

var senderEmail string
var receiverEmail string
var awsRegion string
var awsAccessKeyId string
var awsSecretAccessKey string

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	senderEmail = os.Getenv("SENDER_EMAIL")
	receiverEmail = os.Getenv("RECEIVER_EMAIL")
	awsRegion = os.Getenv("AWS_REGION")
	awsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func main() {

	vars := map[string]string{ //variables that will go to HTML template
		"name":         "evzpav",
		"info":         "123456",
		"supportEmail": senderEmail,
	}

	s := email.NewClient(awsRegion, awsAccessKeyId, awsSecretAccessKey)

	//email based on HTML template
	var emailData = email.EmailData{
		From:         senderEmail,
		To:           []string{receiverEmail},
		ReplyTo:      []string{"noreply@domain.com"},
		Subject:      "My simple email with parsed HTML",
		TemplateName: "email_template.html",
		TemplateVars: vars,
	}
	err := s.Send(emailData)

	if err != nil {
		log.Println(err)
	}
}
