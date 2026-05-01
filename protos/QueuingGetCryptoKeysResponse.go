package protos

type QueuingGetCryptoKeysResponse struct {
	ResponsePacket
	EncryptedKey              string `json:",omitempty,omitzero"`
	SignedKey                 string `json:",omitempty,omitzero"`
	EncryptedIV               string `json:",omitempty,omitzero"`
	SignedIV                  string `json:",omitempty,omitzero"`
	EncryptedSqlCipherKey     string `json:",omitempty,omitzero"`
	EncryptedSqlCipherLicense string `json:",omitempty,omitzero"`
}
