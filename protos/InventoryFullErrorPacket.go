package protos

type InventoryFullErrorPacket struct {
	ResponsePacket
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
