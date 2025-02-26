package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ARMAAN199/Go_EcomApi/config"
	"github.com/ARMAAN199/Go_EcomApi/database"
	"github.com/ARMAAN199/Go_EcomApi/redis"
	router "github.com/ARMAAN199/Go_EcomApi/routers"
	"github.com/ARMAAN199/Go_EcomApi/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Before Server Start")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8088"
	}

	cfg := config.InitConfig()
	redisCfg := config.InitRedisConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	redisClient, err := redis.NewRedisClient(ctx, redisCfg)
	if err != nil {
		log.Fatal(err)
	}

	redisStore := redis.NewDBRedisStore(redisClient)

	err = http.ListenAndServe(PORT, *router.ReturnRouter(db, &redisStore))
	if err != nil {
		log.Fatal(err)
	}
}

func sendMessage() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func recieveMessage() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
