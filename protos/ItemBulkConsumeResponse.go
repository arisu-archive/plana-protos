package protos

type ItemBulkConsumeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UsedItemDB ItemDB `json:",omitempty,omitzero"`
	ParcelInfosInMailBox []ParcelInfo `json:",omitempty,omitzero"`
}
