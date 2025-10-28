package protos

type MomoTalkReadResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MomoTalkOutLineDB MomoTalkOutLineDB `json:",omitempty,omitzero"`
	MomoTalkChoiceDBs []MomoTalkChoiceDB `json:",omitempty,omitzero"`
}
