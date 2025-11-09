package protos

type BillingTransactionStartByYostarRequest struct {
	RequestPacket
	ShopCashId                  int64                        `json:",omitempty,omitzero"`
	ShopCashProductSelectionDBs []ShopCashProductSelectionDB `json:",omitempty,omitzero"`
	VirtualPayment              bool                         `json:",omitempty,omitzero"`
}
