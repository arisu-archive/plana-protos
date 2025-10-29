package protos

type MomoTalkMessageListResponse struct {
	ResponsePacket
	MomoTalkOutLineDB MomoTalkOutLineDB `json:",omitempty,omitzero"`
	MomoTalkChoiceDBs []MomoTalkChoiceDB `json:",omitempty,omitzero"`
}
