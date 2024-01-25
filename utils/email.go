package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)


func SendEmail(toEmail, subject, body string) error {
	emailHost := os.Getenv("EMAIL_HOST")
	emailUser := os.Getenv("EMAIL_USER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPortStr := os.Getenv("EMAIL_PORT")

	// Convert emailPortStr to an integer
	emailPort, err := strconv.Atoi(emailPortStr)
	if err != nil {
		log.Printf("Invalid EMAIL_PORT: %s", emailPortStr)
		return fmt.Errorf("invalid EMAIL_PORT")
	}

	// Set up email configuration
	email := gomail.NewMessage()
	email.SetHeader("From", emailUser)
	email.SetHeader("To", toEmail)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body) // Use "text/html" MIME type for HTML body

	// Set up email sender configuration
	dialer := gomail.NewDialer(emailHost, emailPort, emailUser, emailPassword)

	// Send the email
	if err := dialer.DialAndSend(email); err != nil {
		log.Printf("Error sending email: %v", err)
		return fmt.Errorf("failed to send email")
	}

	log.Println("Email sent successfully!")
	return nil
}


// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/hibiken/asynq"
// )

// // EmailTask represents the task to send an email.
// type EmailTask struct {
// 	To      string
// 	Subject string
// 	Body    string
// }

// // ProcessEmailTask processes the email task.
// func ProcessEmailTask(ctx context.Context, task *asynq.Task) error {
// 	emailTask, ok := task.Payload().(EmailTask)
// 	if !ok {
// 		return fmt.Errorf("invalid task payload type")
// 	}

// 	// Logic to send the email (replace this with your actual email sending logic)
// 	fmt.Printf("Sending email to: %s\nSubject: %s\nBody: %s\n",
// 		emailTask.To, emailTask.Subject, emailTask.Body)

// 	return nil
// }

// func main() {
// 	// Initialize the asynq server.
// 	server := asynq.NewServer(
// 		asynq.RedisClientOpt{Addr: "localhost:6379", DB: 0},
// 		asynq.Config{Concurrency: 10},
// 	)

// 	// Start the server.
// 	if err := server.Run(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create an asynq client.
// 	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379", DB: 0})

// 	// Define an email task.
// 	emailTask := EmailTask{
// 		To:      "recipient@example.com",
// 		Subject: "Hello from Asynq",
// 		Body:    "This is a test email.",
// 	}

// 	// Enqueue the email task.
// 	task := asynq.NewTask("send-email", emailTask)
// 	if _, err := client.Enqueue(task); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Wait for a moment to allow the worker to process the task.
// 	time.Sleep(2 * time.Second)

// 	// Shutdown the server.
// 	if err := server.Shutdown(context.Background()); err != nil {
// 		log.Fatal(err)
// 	}
// }
