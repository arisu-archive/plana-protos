package protos

type MiniGameDreamMakerResetResponse struct {
	ResponsePacket
	InfoDB       MiniGameDreamMakerInfoDB
	ParameterDBs []MiniGameDreamMakerParameterDB
}
