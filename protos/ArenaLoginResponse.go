package protos

type ArenaLoginResponse struct {
	ResponsePacket
	ArenaPlayerInfoDB ArenaPlayerInfoDB `json:",omitempty,omitzero"`
}
