package protos

type AccountPassCheckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	DevId string `json:",omitempty,omitzero"`
	OnlyAccountId bool `json:",omitempty,omitzero"`
	ClientGeneratedKey string `json:",omitempty,omitzero"`
	ClientGeneratedIV string `json:",omitempty,omitzero"`
}
