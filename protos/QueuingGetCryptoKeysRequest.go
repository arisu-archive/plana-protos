package protos

type QueuingGetCryptoKeysRequest struct {
	RequestPacket
	ClientGeneratedKey string
	ClientGeneratedIV  string
}
