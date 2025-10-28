package protos

type BillingPurchaseListByYostarRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
