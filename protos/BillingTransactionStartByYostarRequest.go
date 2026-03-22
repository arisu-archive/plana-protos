package protos

type BillingTransactionStartByYostarRequest struct {
	RequestPacket
	ShopCashId                  int64 `json:",omitempty,omitzero"`
	ShopCashProductSelectionDBs []ShopCashProductSelectionDB
	VirtualPayment              bool `json:",omitempty,omitzero"`
}
