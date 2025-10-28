package protos

type MiniGameDreamMakerNewGameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	InfoDB MiniGameDreamMakerInfoDB `json:",omitempty,omitzero"`
	ParameterDBs []MiniGameDreamMakerParameterDB `json:",omitempty,omitzero"`
}
