package api

import (
	"encoding/json"

	"fmt"

	"github.com/iotaledger/iota.go/api"
)

var endpoint = "http://node.iotadev.org:14265"

type iotaAPI struct {
}

// IotaAPI ...
type IotaAPI interface {
	IotaSender
	IotaReceiver
	IotaAddressCreator
	IotaBalanceChecker
	IotaDataSender
}

// IotaSender ...
type IotaSender interface {
	send(api iotaAPI) error
}

// IotaReceiver ...
type IotaReceiver interface {
	receive(api iotaAPI) (boolean error)
}

// IotaAddressCreator ...
type IotaAddressCreator interface {
	createAddress(api iotaAPI) (string error)
}

// IotaBalanceChecker ...
type IotaBalanceChecker interface {
	checkBalance(api iotaAPI) (string error)
}

// IotaDataSender ...
type IotaDataSender interface {
	send(api iotaAPI) error
}

func (iota *iotaAPI) send(api iotaAPI) error {

}

// API connects to API of a IRI node on the Tangle network
func API() {

	// compose a new API instance
	iotaAPI, err := api.ComposeAPI(api.HTTPClientSettings{URI: endpoint})
	must(err)

	nodeInfo, err := iotaAPI.GetNodeInfo()
	must(err)

	pretty, err := json.MarshalIndent(nodeInfo, "", "    ")
	must(err)

	fmt.Println("latest milestone index:", nodeInfo.LatestMilestoneIndex)
	fmt.Println(string(pretty))
	return iotaAPI
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
