package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
)

type pulsarConnection struct {
	dsn    string
	Client pulsar.Client
}

// New returns a pointer to a pulsarConnection with an established db session, or an error.
func newPulsarConnection(dsn string, timeoutSeconds time.Duration) (*pulsarConnection, error) {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               dsn,
		OperationTimeout:  timeoutSeconds * time.Second,
		ConnectionTimeout: timeoutSeconds * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("could not create pulsar client, err = %w", err)
	}

	conn := pulsarConnection{
		dsn:    dsn,
		Client: client,
	}
	if err := conn.Check(); err != nil {
		return nil, fmt.Errorf("pulsar connection check failed, err = %w", err)
	}

	return &conn, nil
}

// connectPulsar returns a Pulsar connection or an error.
func connectPulsar(dsn string, timeoutSeconds time.Duration) (*pulsarConnection, error) {
	return newPulsarConnection(dsn, timeoutSeconds)
}

// Check the client connection by creating a reader
func (c *pulsarConnection) Check() error {

	_, err := c.Client.CreateProducer(pulsar.ProducerOptions{
		Topic: uuid.New().String(),
	})
	return err
}

// Publish will send a message to pulsar on the specified topic.
func (c *pulsarConnection) Publish(topic string, payload []byte) error {

	producer, err := c.Client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return fmt.Errorf("could not create producer, err = %w", err)
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: payload,
	})
	if err != nil {
		return fmt.Errorf("could not publish message, err = %w", err)
	}
	producer.Close()

	return nil
}

// PublishJSON is a convenience method that will marshal a JSON payload
// into a []byte and the call Publish().
func (c *pulsarConnection) PublishJSON(topic string, jsonPayload string) error {
	xb, err := json.Marshal(jsonPayload)
	if err != nil {
		return fmt.Errorf("could not marshal json, err = %s", err)
	}
	return c.Publish(topic, xb)
}

func (c *pulsarConnection) Close() {
	c.Client.Close()
}
