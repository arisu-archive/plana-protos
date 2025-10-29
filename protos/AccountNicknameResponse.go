package protos

type AccountNicknameResponse struct {
	ResponsePacket
	AccountDB AccountDB `json:",omitempty,omitzero"`
}
