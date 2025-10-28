package protos

type MiniGameDreamMakerNewGameRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Multiplier int64 `json:",omitempty,omitzero"`
}
