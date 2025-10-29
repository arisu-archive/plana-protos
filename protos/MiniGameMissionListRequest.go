package protos

type MiniGameMissionListRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
