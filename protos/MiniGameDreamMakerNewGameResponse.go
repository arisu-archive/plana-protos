package protos

type MiniGameDreamMakerNewGameResponse struct {
	ResponsePacket
	InfoDB       MiniGameDreamMakerInfoDB
	ParameterDBs []MiniGameDreamMakerParameterDB
}
