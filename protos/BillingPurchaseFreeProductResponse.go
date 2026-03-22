package protos

type BillingPurchaseFreeProductResponse struct {
	ResponsePacket
	ParcelResult    ParcelResultDB
	MailDB          MailDB
	PurchaseProduct PurchaseCountDB
}
