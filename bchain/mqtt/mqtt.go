package mqtt

import (
	"fmt"
	"log"
	"time"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

const (
	broker = "iot.eclipse.org:1883"
)

var (
	cli *client.Client
)

func init() {

}

// Publish publishes message on given topic
func Publish(topic, message string) (err error) {
	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			log.Println(err)
		},
	})
	defer cli.Terminate()

	// Connect to the MQTT Server.
	err = cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  broker,
		ClientID: []byte("wty"),
	})
	if err != nil {
		return err
	}

	// Subscribe to topics.
	err = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte("centre/blocks"),
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: func(topicName, message []byte) {
					log.Println("received", string(topicName), string(message))
				},
			},
			&client.SubReq{
				TopicFilter: []byte("foo"),
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: func(topicName, message []byte) {
					fmt.Println("received", string(topicName), string(message))
				},
			},
			&client.SubReq{
				TopicFilter: []byte("minors/connected"),
				QoS:         mqtt.QoS1,
				Handler: func(topicName, message []byte) {
					log.Println("received", string(topicName), string(message))
				},
			},
		},
	})
	// Publish a message.
	err = cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(topic),
		Message:   []byte(message),
	})
	if err != nil {
		return err
	}

	time.Sleep(time.Second )
	// Disconnect the Network Connection.
	if err := cli.Disconnect(); err != nil {
		return err
	}

	return
}
