package protos

type EventContentPortalRequest struct {
	RequestPacket
	EventContentId  int64 `json:",omitempty,omitzero"`
	StageUniqueId   int64 `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
}
