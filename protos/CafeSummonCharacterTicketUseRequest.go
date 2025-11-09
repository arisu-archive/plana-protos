package protos

type CafeSummonCharacterTicketUseRequest struct {
	RequestPacket
	CafeDBId          int64 `json:",omitempty,omitzero"`
	CharacterServerId int64 `json:",omitempty,omitzero"`
}
