package protos

type MomoTalkFavorScheduleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScheduleId int64 `json:",omitempty,omitzero"`
}
