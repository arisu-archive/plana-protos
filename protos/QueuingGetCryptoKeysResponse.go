package protos

type QueuingGetCryptoKeysResponse struct {
	ResponsePacket
	EncryptedKey              string
	SignedKey                 string
	EncryptedIV               string
	SignedIV                  string
	EncryptedSqlCipherKey     string
	EncryptedSqlCipherLicense string
}
