package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	kafka "github.com/kveriz/kfproxy/internal/client"
	server "github.com/kveriz/kfproxy/internal/server"
)

var (
	port string
)

func main() {
	flag.StringVar(&port, "port", ":8080", "Port to serve requests")
	flag.Parse()

	conf, ok := os.LookupEnv("PRODUCER_CONFIG")

	if !ok {
		conf = fmt.Sprintf("%s/producer.properties", os.Getenv("HOME"))
		log.Print("try to use default config's location", conf)
	}

	propFile, err := os.Open(conf)

	if err != nil {
		log.Fatal(err)
	}

	defer propFile.Close()

	scanner := bufio.NewScanner(propFile)

	opts := make(map[string]interface{})

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		var prop []string

		if !strings.HasPrefix(scanner.Text(), "#") {
			prop = strings.Split(scanner.Text(), "=")
		}
		key := prop[0]
		value := prop[1]

		opts[key] = value
	}

	log.Print("reading of config file is complete")

	s := server.New(port)
	cfg := kafka.NewConfig(opts)
	kf := kafka.New(*cfg)

	s.Serve(kf.ServeHTTP)
}
