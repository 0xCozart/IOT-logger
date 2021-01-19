package transactions

import (
	"fmt"
	"time"

	"github.com/AlanVegaDecentralize/IOT-logger/utils"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/trinary"
)

// TrxData is the basic data needed to execute a transaction from the user
type TrxData struct {
	Address string
	Tag     string
	Value   uint64
	Message string
}

// CreateTx accepts TrxData and returns bundle for transfer
func CreateTx(data TrxData) {
	addressTrytes, err := trinary.NewTrytes(data.Address)
	utils.Must("address", err)
	tagTrytes, err := trinary.NewTrytes(data.Tag)
	utils.Must("tag", err)
	messageTrytes, err := trinary.NewTrytes(data.Message)
	utils.Must("message", err)

	ts := uint64(time.Now().UnixNano() / int64(time.Second))
	transfers := bundle.Transfers{{
		Address: addressTrytes,
		Tag:     tagTrytes,
		Value:   data.Value,
		Message: messageTrytes,
	},
	}
	fmt.Println(transfers, ts)
}
