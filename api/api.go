package api

import (
	"github.com/iotaledger/iota.go/api"

	"fmt"
)

var endpoint = "http://node.iotadev.org:14265"

// API connects to API of a IRI node on the Tangle network
func API() {

	// compose a new API instance
	api, err := api.ComposeAPI(api.HTTPClientSettings{URI: endpoint})
	must(err)

	nodeInfo, err := api.GetNodeInfo()
	must(err)

	fmt.Println("latest milestone index:", nodeInfo.LatestMilestoneIndex)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
