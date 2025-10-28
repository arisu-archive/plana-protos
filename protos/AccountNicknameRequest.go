package protos

type AccountNicknameRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Nickname string `json:",omitempty,omitzero"`
}
