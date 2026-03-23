package protos

type ItemBulkConsumeResponse struct {
	ResponsePacket
	UsedItemDB           *ItemDB `json:",omitempty,omitzero"`
	ParcelInfosInMailBox []ParcelInfo
}
