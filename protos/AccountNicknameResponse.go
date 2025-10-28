package protos

type AccountNicknameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountDB AccountDB `json:",omitempty,omitzero"`
}
