package protos

type RaidSearchRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SecretCode string `json:",omitempty,omitzero"`
	Tags []string `json:",omitempty,omitzero"`
}
