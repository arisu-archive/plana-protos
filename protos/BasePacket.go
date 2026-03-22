package protos

type BasePacket struct {
	SessionKey SessionKey
	Protocol   Protocol `json:",omitempty,omitzero"`
	AccountId  int64    `json:",omitempty,omitzero"`
}
