package protos

type MiniGameTableBoardMoveThemaRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
