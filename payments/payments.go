package payments

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"
)

type Data struct {
	Clientid  int       `json:"clientid"`
	Siteid    int       `json:"siteid"`
	Provider  string    `json:"provider"`
	MachineId string    `json:"machine_id,omitempty"`
	VRM       string    `json:"vrm"`
	DateFrom  string    `json:"date_from"`
	DateTo    string    `json:"date_to"`
	PaymentID string    `json:"payment_id,omitempty"`
	RawData   string    `json:"raw_data"`
	Received  time.Time `json:"received"`
}

func (d *Data) Save(pstopic string) error {

	if d.Siteid == 0 {
		return errors.New("missing siteid")
	}

	if d.Clientid == 0 {
		return errors.New("missing siteid")
	}

	if len(d.VRM) == 0 {
		return errors.New("missing vehicle information")
	}

	if d.Received.Year() == 0 {
		d.Received = time.Now()
	}

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "notanotherapp")

	defer client.Close()

	if err != nil {
		return err
	}

	topic := client.Topic(pstopic)

	attrib := make(map[string]string)

	attrib["siteid"] = strconv.Itoa(d.Siteid)

	payload, err := json.Marshal(&d)

	if err != nil {
		log.Println(err)
		return err
	}

	msg := &pubsub.Message{
		Data: payload,
		//Attributes: attrib,
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		return err
	}

	topic.Stop()

	return nil

}
