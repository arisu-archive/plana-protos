package protos

type InventoryFullErrorPacket struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
