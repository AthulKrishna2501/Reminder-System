package main

import (
	"log"
	"remainder-system/Remainder-Service/db"
	rabbitmq "remainder-system/Remainder-Service/rabbit-mq"
	"remainder-system/Remainder-Service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := rabbitmq.Init(); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ %v", err)
	}

	defer rabbitmq.Close()
	db.InitDB()

	r := gin.Default()

	r.POST("/create-reminder", routes.CreateReminder)

	log.Println("Starting server on port 5000...")
	r.Run(":5000")
}
