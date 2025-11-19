package protos

type MomoTalkReadResponse struct {
	ResponsePacket
	MomoTalkOutLineDB MomoTalkOutLineDB  `json:",omitempty,omitzero"`
	MomoTalkChoiceDBs []MomoTalkChoiceDB `json:",omitempty,omitzero"`
}
