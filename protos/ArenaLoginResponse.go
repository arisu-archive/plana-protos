package protos

type ArenaLoginResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ArenaPlayerInfoDB ArenaPlayerInfoDB `json:",omitempty,omitzero"`
}
