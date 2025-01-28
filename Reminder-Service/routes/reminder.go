package routes

import (
	"net/http"
	"remainder-system/Remainder-Service/db"
	"remainder-system/Remainder-Service/models"
	rabbitmq "remainder-system/Remainder-Service/rabbit-mq"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func CreateReminder(c *gin.Context) {
	var reminder models.Reminder

	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newReminder := models.Reminder{
		UserEmail:    reminder.UserEmail,
		Title:        reminder.Title,
		Description:  reminder.Description,
		ReminderTime: reminder.ReminderTime,
	}

	if err := db.Db.Create(&newReminder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	if rabbitmq.Channel == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "RabbitMQ channel is not open"})
		return
	}

	err := rabbitmq.Channel.Publish(
		"",
		"reminder_queue",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(newReminder.UserEmail + "|" + newReminder.Title),
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish message to RabbitMQ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reminder created successfully"})
}
