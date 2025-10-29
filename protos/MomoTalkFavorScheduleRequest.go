package protos

type MomoTalkFavorScheduleRequest struct {
	RequestPacket
	ScheduleId int64 `json:",omitempty,omitzero"`
}
