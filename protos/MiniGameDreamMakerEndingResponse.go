package protos

type MiniGameDreamMakerEndingResponse struct {
	ResponsePacket
	InfoDB         *MiniGameDreamMakerInfoDB `json:",omitempty,omitzero"`
	ParameterDBs   []*MiniGameDreamMakerParameterDB
	EndingDB       *MiniGameDreamMakerEndingDB `json:",omitempty,omitzero"`
	ParcelResultDB *ParcelResultDB             `json:",omitempty,omitzero"`
}
