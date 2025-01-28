package main

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel %v", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"reminder_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare queue %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to cosume queue %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			SendMail(string(d.Body))
		}
	}()

	log.Print("Waiting for messages....")
	<-forever
}

func SendMail(message string) {
	parts := SplitMessage(message)
	if len(parts) < 2 {
		log.Println("Invalid message format")
		return
	}
	userEmail := parts[0]
	reminderTitle := parts[1]

	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	log.Print(from, password)
	to := []string{userEmail}
	subject := "Subject: Reminder: " + reminderTitle + "\n"
	body := "Hi there,\n\nThis is your reminder for: " + reminderTitle + ".\n\nBest regards,\nReminder System"
	msg := []byte(subject + "\n" + body)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from,
		to,
		msg,
	)
	if err != nil {
		log.Printf("Failed to send email: %s", err)
		return
	}

	log.Println("Email sent successfully!")
}

func SplitMessage(message string) []string {
	return strings.Split(message, "|")
}
