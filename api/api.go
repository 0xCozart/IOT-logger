package api

import (
	"encoding/json"

	"fmt"

	"github.com/iotaledger/iota.go/api"
)

const (
	seed     = "CDVOPWJ9BBM9OHNRJXKINKXYANHKZIFBDQLWQBRIIFAJCEZI9YWWN9IZPKNVFERRBDRMTWPDLOAKWQCZL"
	endpoint = "http://node.iotadev.org:14265"
)

type iotaAPI struct{}

// IotaAPI ...
type IotaAPI interface {
	send() error
	// IotaReceiver
	// IotaAddressCreator
	// IotaBalanceChecker
	// IotaDataSender
}

// // IotaSender ...
// type IotaSender interface {
// 	send(api iotaAPI) error
// }

// // IotaReceiver ...
// type IotaReceiver interface {
// 	receive(api iotaAPI) (boolean error)
// }

// // IotaAddressCreator ...
// type IotaAddressCreator interface {
// 	createAddress(api iotaAPI) (string error)
// }

// // IotaBalanceChecker ...
// type IotaBalanceChecker interface {
// 	checkBalance(api iotaAPI) (string error)
// }

// // IotaDataSender ...
// type IotaDataSender interface {
// 	send(api iotaAPI) error
// }

// type IotaTransactionCreator interface {
// }

// SendTx ...
func (api *iotaAPI) SendTx(txInfo) error {

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
