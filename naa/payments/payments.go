package payments

import "time"

const PAYBYPHONE = "paybyphone"
const TAP2PARK = "tap2park"

type Data struct {
	Clientid  int       `json:"clientid"`
	Siteid    int       `json:"siteid"`
	Provider  string    `json:"provider"`
	MachineId string    `json:"machine_id"`
	VRM       string    `json:"vrm"`
	DateFrom  string    `json:"date_from"`
	DateTo    string    `json:"date_to"`
	PaymentID string    `json:"payment_id"`
	RawData   string    `json:"raw_data"`
	Received  time.Time `json:"received"`
}

func (d *Data) Save() {

}
