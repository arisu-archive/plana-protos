package protos

type MomoTalkOutLineResponse struct {
	ResponsePacket
	MomoTalkOutLineDBs   []MomoTalkOutLineDB `json:",omitempty,omitzero"`
	FavorScheduleRecords map[int64][]int64   `json:",omitempty,omitzero"`
}
