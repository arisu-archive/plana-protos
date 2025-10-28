package protos

type AccountBanErrorPacket struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
	BanReason string `json:",omitempty,omitzero"`
}
