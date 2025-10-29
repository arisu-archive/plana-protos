package protos

type BillingPurchaseFreeProductResponse struct {
	ResponsePacket
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
	MailDB MailDB `json:",omitempty,omitzero"`
	PurchaseProduct PurchaseCountDB `json:",omitempty,omitzero"`
}
