package protos

type EventContentCardShopShuffleRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
