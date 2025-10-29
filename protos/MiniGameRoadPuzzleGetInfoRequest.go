package protos

type MiniGameRoadPuzzleGetInfoRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
