package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

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

	uuid, err := uuid.NewUUID()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	topics := r.Header["Topicname"]

	body := make([]byte, 0)

	defer r.Body.Close()

	body, err = io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	value := make(map[string]interface{})

	err = json.Unmarshal(body, &value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	key := uuid.String()

	if err := c.produce(topics, key, value); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Message was successfully published to topic %v with key %s", topics, key)))
}

func (c *Kafka) produce(t []string, k string, v map[string]any) error {
	value, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	topics := strings.Split(t[0], ",")
	go func(t []string) {
		for i, _ := range t {
			topic := &kafka.TopicPartition{
				Topic:     &t[i],
				Partition: kafka.PartitionAny,
			}

			msg := &kafka.Message{
				TopicPartition: *topic,
				Value:          value,
				Key:            []byte(k),
				Timestamp:      time.Now(),
			}
			err = c.сlient.Produce(msg, nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	}(topics)

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
