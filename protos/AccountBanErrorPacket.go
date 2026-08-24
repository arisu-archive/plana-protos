package protos

type AccountBanErrorPacket struct {
	ResponsePacket
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
	BanReason string
}
