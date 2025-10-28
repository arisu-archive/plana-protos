package protos

type MiniGameDreamMakerEndingResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	InfoDB MiniGameDreamMakerInfoDB `json:",omitempty,omitzero"`
	ParameterDBs []MiniGameDreamMakerParameterDB `json:",omitempty,omitzero"`
	EndingDB MiniGameDreamMakerEndingDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
