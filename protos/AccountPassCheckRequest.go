package protos

type AccountPassCheckRequest struct {
	RequestPacket
	DevId         string
	OnlyAccountId bool `json:",omitempty,omitzero"`
}
