package api

import (
	"encoding/json"

	"github.com/iotaledger/iota.go/api"

	"fmt"
)

var endpoint = "http://node.iotadev.org:14265"

type IotaAPI struct{
	API: api.Api
}

// API connects to API of a IRI node on the Tangle network
func API() IotaAPI {

	// compose a new API instance
	iotaAPI, err := api.ComposeAPI(api.HTTPClientSettings{URI: endpoint})
	must(err)

	nodeInfo, err := iotaAPI.GetNodeInfo()
	must(err)

	pretty, err := json.MarshalIndent(nodeInfo, "", "    ")
	must(err)

	fmt.Println("latest milestone index:", nodeInfo.LatestMilestoneIndex)
	fmt.Println(string(pretty))
	return iotaAPI.api
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
