package protos

type RaidSearchRequest struct {
	RequestPacket
	SecretCode string `json:",omitempty,omitzero"`
	Tags []string `json:",omitempty,omitzero"`
}
