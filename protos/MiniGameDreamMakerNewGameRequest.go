package protos

type MiniGameDreamMakerNewGameRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	Multiplier int64 `json:",omitempty,omitzero"`
}
