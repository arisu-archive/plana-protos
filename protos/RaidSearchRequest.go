package protos

type RaidSearchRequest struct {
	RequestPacket
	SecretCode string
	Tags       []string
}
