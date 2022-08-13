package kafka_pac

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/mohammad-siraj/crud_kafka/entities"
)

type producer struct {
	prod *kafka.Producer
}

type Producer interface {
	ProducerPub(car entities.Car, s string)
	Initialproducer() error
}

func (pro *producer) Initialproducer() error {
	m := entities.Car{
		Model: "this is a",
		Make:  "kafka",
		Year:  "messsage",
	}
	go produce_check(pro.prod)
	//p := producer{prod: pro}
	pro.ProducerPub(m, "reponse")
	return nil
}

func NewProducer() (*producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9091"})
	return &producer{prod: p}, err
}

func produce_check(p *kafka.Producer) {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("\nMessage not delived %v", ev.TopicPartition.Error)
			} else {
				fmt.Printf("\nMessage delived sucessfully\n")
			}
		}
	}
}

func (p producer) ProducerPub(car entities.Car, topicname string) {
	me, _ := json.Marshal(car)
	topic := topicname
	data := []byte(string(me))

	p.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: int32(kafka.PartitionAny)},
		Value: data,
	}, nil)
	//p.prod.Flush(15)
}
