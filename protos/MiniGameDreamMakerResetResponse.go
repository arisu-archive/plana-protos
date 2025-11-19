package protos

type MiniGameDreamMakerResetResponse struct {
	ResponsePacket
	InfoDB       MiniGameDreamMakerInfoDB        `json:",omitempty,omitzero"`
	ParameterDBs []MiniGameDreamMakerParameterDB `json:",omitempty,omitzero"`
}
