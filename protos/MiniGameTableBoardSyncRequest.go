package protos

type MiniGameTableBoardSyncRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
