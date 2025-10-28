package protos

type BillingPurchaseFreeProductResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
	MailDB MailDB `json:",omitempty,omitzero"`
	PurchaseProduct PurchaseCountDB `json:",omitempty,omitzero"`
}
