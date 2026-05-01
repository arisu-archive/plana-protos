package protos

type AccountNicknameRequest struct {
	RequestPacket
	Nickname string `json:",omitempty,omitzero"`
}
