package protos

type MiniGameTableBoardMoveRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	Steps []HexLocation `json:",omitempty,omitzero"`
}
