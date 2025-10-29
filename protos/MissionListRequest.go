package protos

type MissionListRequest struct {
	RequestPacket
	EventContentId *int64 `json:",omitempty,omitzero"`
}
