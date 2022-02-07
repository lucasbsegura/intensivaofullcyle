package main

import (
	kafka "course/infra/kafka"
	repository2 "course/infra/repository"
	usecase2 "course/usecase"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/fullcycle")
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository2.CourseMySqlRepository{Db: db}
	usecase := usecase2.CreateCourse{Repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
	}
	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	consumer.Consume(msgChan)

	for msg := range msgChan {
		var input usecase2.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)
		output, err := usecase.Execute(input)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println(output)
		}
	}
}
