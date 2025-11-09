package protos

type BasePacket struct {
	SessionKey SessionKey `json:",omitempty,omitzero"`
	Protocol   Protocol   `json:",omitempty,omitzero"`
	AccountId  int64      `json:",omitempty,omitzero"`
}
