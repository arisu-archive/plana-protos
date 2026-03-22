package protos

type ItemBulkConsumeResponse struct {
	ResponsePacket
	UsedItemDB           ItemDB
	ParcelInfosInMailBox []ParcelInfo
}
