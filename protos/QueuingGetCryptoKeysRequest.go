package protos

type QueuingGetCryptoKeysRequest struct {
	RequestPacket
	ClientGeneratedKey string `json:",omitempty,omitzero"`
	ClientGeneratedIV  string `json:",omitempty,omitzero"`
}
