package main
import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"context"
	"log"
)


type EmailData struct {
	From           string
	To             []string
	CC             []string
	BCC            []string
	ReplyTo        []string
	Subject        string
	Text           string
	HTML           string
	TemplateName   string
	TemplateVars   interface{}
	AttachFiles    []string
	ConfigSet      string
	BaseLayoutPath string
}

func main(){
	fmt.Println("test")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	ses_client:= ses.NewFromConfig(cfg)
	ss:=types.Message{}
	fmt.Println(ss)
	fmt.Println(ses_client)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("fletch-widget-asset"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}


func createInput(mail *EmailData) (*ses.SendEmailInput, error) {
	if mail.HTML == "" {
		mail.HTML = mail.Text
	}
	msg := &types.Message{
		Subject: &types.Content{
			Charset: aws.String("utf-8"),
			Data:    &mail.Subject,
		},
		Body: &types.Body{
			Html: &types.Content{
				Charset: aws.String("utf-8"),
				Data:    &mail.HTML,
			},
			Text: &types.Content{
				Charset: aws.String("utf-8"),
				Data:    &mail.Text,
			},
		},
	}

	dest := &types.Destination{
		ToAddresses:  aws.StringSlice(mail.To),
		CcAddresses:  aws.StringSlice(mail.CC),
		BccAddresses: aws.StringSlice(mail.BCC),
	}

	return &types.SendEmailInput{
		Source:               &mail.From,
		Destination:          dest,
		Message:              msg,
		ReplyToAddresses:     aws.StringSlice(mail.ReplyTo),
		ConfigurationSetName: aws.String(mail.ConfigSet),
	}, nil

}