package protos

type EventContentMapMoveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
	DestPosition HexLocation `json:",omitempty,omitzero"`
}
