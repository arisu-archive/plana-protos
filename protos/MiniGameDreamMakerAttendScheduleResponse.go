package protos

type MiniGameDreamMakerAttendScheduleResponse struct {
	ResponsePacket
	InfoDB                    *MiniGameDreamMakerInfoDB `json:",omitempty,omitzero"`
	ParameterDBs              []*MiniGameDreamMakerParameterDB
	ParcelResultDB            *ParcelResultDB `json:",omitempty,omitzero"`
	ScheduleResultId          int64           `json:",omitempty,omitzero"`
	EventContentCollectionDBs []*EventContentCollectionDB
}
