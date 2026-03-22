package protos

type AccountPassCheckResponse struct {
	ResponsePacket
	EncryptedKey string
	SignedKey    string
	EncryptedIV  string
	SignedIV     string
}
