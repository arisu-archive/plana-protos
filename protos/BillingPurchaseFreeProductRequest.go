package protos

type BillingPurchaseFreeProductRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopCashId int64 `json:",omitempty,omitzero"`
}
