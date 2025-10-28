package protos

type MiniGameTableBoardMoveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Steps []HexLocation `json:",omitempty,omitzero"`
}
