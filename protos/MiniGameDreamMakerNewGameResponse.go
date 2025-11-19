package protos

type MiniGameDreamMakerNewGameResponse struct {
	ResponsePacket
	InfoDB       MiniGameDreamMakerInfoDB        `json:",omitempty,omitzero"`
	ParameterDBs []MiniGameDreamMakerParameterDB `json:",omitempty,omitzero"`
}
