package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/google/uuid"
)

type Kafka struct {
	сlient *kafka.Producer
}

func New(c KFConfig) *Kafka {
	p, err := kafka.NewProducer(c.CFG)

	if err != nil {
		log.Fatal(err)
	}

	return &Kafka{
		сlient: p,
	}
}

func (c *Kafka) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	key, err := uuid.NewUUID()

	if err != nil {
		return
	}

	topic := r.Header["Topicname"][0]

	_ = topic
	fmt.Println(topic)

	body := make([]byte, 0)

	defer r.Body.Close()

	body, err = io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	value := make(map[string]interface{})

	err = json.Unmarshal(body, &value)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)

	w.Write([]byte(key.String()))
	w.Write([]byte(topic))

	err = c.produce(topic, key.String(), value)

}

func (c *Kafka) produce(t, k string, v map[string]any) error {
	topic := &kafka.TopicPartition{
		Topic:     &t,
		Partition: kafka.PartitionAny,
	}

	value, err := json.Marshal(v)

	if err != nil {
		return err
	}

	fmt.Println(string(value))

	msg := &kafka.Message{
		TopicPartition: *topic,
		Value:          value,
		Key:            []byte(k),
		Timestamp:      time.Now(),
	}

	// fmt.Printf("%#v", msg)

	err = c.сlient.Produce(msg, nil)

	if err != nil {
		return err
	}

	go func() {
		for e := range c.сlient.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	return nil
}
