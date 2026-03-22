package protos

type MomoTalkReadResponse struct {
	ResponsePacket
	MomoTalkOutLineDB MomoTalkOutLineDB
	MomoTalkChoiceDBs []MomoTalkChoiceDB
}
