package protos

type MomoTalkMessageListResponse struct {
	ResponsePacket
	MomoTalkOutLineDB MomoTalkOutLineDB
	MomoTalkChoiceDBs []MomoTalkChoiceDB
}
