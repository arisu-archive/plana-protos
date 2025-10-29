package protos

type MiniGameTableBoardResurrectRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
