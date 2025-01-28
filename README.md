# Personalized Email Reminder System

A microservices-based project that allows users to create reminders and sends them personalized email notifications at the appropriate time. This project uses RabbitMQ for messaging, SendGrid (or SMTP) for email delivery, and is implemented in Golang.

## Features
- **Reminder Service**: Create and store reminders in a PostgreSQL database.
- **Notification Service**: Listens to RabbitMQ for new reminders and sends email notifications.
- **Message Queue**: RabbitMQ facilitates communication between services.

---

## Tech Stack
- **Backend**: Golang
- **Database**: PostgreSQL
- **Message Queue**: RabbitMQ
- **Email Delivery**: SendGrid (or SMTP)
- **Containerization**: Docker
- **Framework**: Gin (for the Reminder Service API)

---

## Prerequisites
- Install [Go](https://golang.org/)
- Install [Docker](https://www.docker.com/) (to run RabbitMQ)
- Install PostgreSQL

---

## Setup Guide

### 1. Clone the Repository
```bash
git clone <repository-url>
cd Reminder-System
```

### 2. Set Up RabbitMQ

#### Option A: Using Docker (Recommended)
Run the following command to start a RabbitMQ container:
```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```
Access the RabbitMQ Management Console at [http://localhost:15672/](http://localhost:15672/):
- Username: `guest`
- Password: `guest`

#### Option B: Install RabbitMQ Locally
Follow the [official RabbitMQ installation guide](https://www.rabbitmq.com/download.html).

---

### 3. Set Up PostgreSQL
- Install PostgreSQL if not already installed.
- Create a database:
```sql
CREATE DATABASE reminder_system;
```
- Configure the database connection in the `.env` file.

---

### 4. Configure Environment Variables
Create a `.env` file in the project root directory with the following:
```env
# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@localhost:5672/

# PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=reminder_system

# Email Credentials
EMAIL=your-email@example.com
PASSWORD=your-app-password
```

---

### 5. Run the Services

#### Reminder Service
Start the Reminder Service:
```bash
cd Remainder-Service
go run main.go
```

#### Notification Service
Start the Notification Service:
```bash
cd Notification-Service
go run main.go
```

---

## API Endpoints

### Reminder Service
#### Create a Reminder
- **Endpoint**: `POST /reminders`
- **Body**:
```json
{
  "user_email": "user@example.com",
  "title": "Test Reminder",
  "description": "This is a test reminder.",
  "reminder_time": "2025-01-28T14:30:00Z"
}
```
- **Response**:
```json
{
  "message": "Reminder created successfully"
}
```

---

## Folder Structure
```
Reminder-System/
├── Remainder-Service/
│   ├── db/          # Database setup and models
│   ├── routes/      # API routes for reminders
│   ├── rabbit-mq/   # RabbitMQ connection setup
│   └── main.go      # Main file for the Reminder Service
├── Notification-Service/
│   ├── main.go      # Main file for the Notification Service
├── .env.example     # Example environment variables
└── README.md        # Project documentation
```

---

## Troubleshooting

### Common Issues
1. **RabbitMQ Connection Error**: Ensure RabbitMQ is running on `localhost:5672`. Use Docker logs to debug:
   ```bash
   docker logs rabbitmq
   ```

2. **Environment Variables Not Loaded**: Check if `.env` exists and is in the correct format. Use `godotenv` to load it.

3. **Email Not Sent**: Verify your SMTP credentials and ensure less secure app access is enabled for Gmail accounts.

---

## Future Enhancements
- Add user authentication.
- Implement scheduling for reminders.
- Support multiple email providers.

---

## License
This project is licensed under the MIT License.

