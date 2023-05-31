package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func sendToKindle(mobiFileName string) error {
	from := mail.NewEmail("hackernews", sendFromEmail)
	subject := "Your Daily HackerNews Digest"
	to := mail.NewEmail("kindle", kindleEmail)

	// Read the MOBI file and encode it as base64.
	mobiData, err := ioutil.ReadFile(mobiFileName)
	if err != nil {
		log.Println(err)
		return err
	}
	encodedMobiData := base64.StdEncoding.EncodeToString(mobiData)

	// Create a new attachment.
	attachment := mail.NewAttachment()
	attachment.SetContent(encodedMobiData)
	attachment.SetType("application/x-mobipocket-ebook")
	attachment.SetFilename(mobiFileName)
	attachment.SetDisposition("attachment")
	attachment.SetContentID("Content ID")

	// Add the attachment to the email message.
	message := mail.NewSingleEmail(from, subject, to, "", "")
	message.AddAttachment(attachment)

	plainTextContent := mail.NewContent("text/plain", "This is the email content")
	message.AddContent(plainTextContent)

	// Send the email.
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return err
	}
}
