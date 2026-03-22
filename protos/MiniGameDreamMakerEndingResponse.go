package protos

type MiniGameDreamMakerEndingResponse struct {
	ResponsePacket
	InfoDB         MiniGameDreamMakerInfoDB
	ParameterDBs   []MiniGameDreamMakerParameterDB
	EndingDB       MiniGameDreamMakerEndingDB
	ParcelResultDB ParcelResultDB
}
