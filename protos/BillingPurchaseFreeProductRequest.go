package protos

type BillingPurchaseFreeProductRequest struct {
	RequestPacket
	ShopCashId int64 `json:",omitempty,omitzero"`
}
