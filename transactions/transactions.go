package transactions

// TrxData is the basic data needed to execute a transaction from the user
type TrxData struct {
	Address string
	Value   uint64
	Message
}
