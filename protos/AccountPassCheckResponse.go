package protos

type AccountPassCheckResponse struct {
	ResponsePacket
	EncryptedKey string `json:",omitempty,omitzero"`
	SignedKey string `json:",omitempty,omitzero"`
	EncryptedIV string `json:",omitempty,omitzero"`
	SignedIV string `json:",omitempty,omitzero"`
}
