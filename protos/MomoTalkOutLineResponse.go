package protos

type MomoTalkOutLineResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MomoTalkOutLineDBs []MomoTalkOutLineDB `json:",omitempty,omitzero"`
	FavorScheduleRecords map[int64][]int64 `json:",omitempty,omitzero"`
}
