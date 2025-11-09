package protos

type MiniGameDreamMakerAttendScheduleResponse struct {
	ResponsePacket
	InfoDB                    MiniGameDreamMakerInfoDB        `json:",omitempty,omitzero"`
	ParameterDBs              []MiniGameDreamMakerParameterDB `json:",omitempty,omitzero"`
	ParcelResultDB            ParcelResultDB                  `json:",omitempty,omitzero"`
	ScheduleResultId          int64                           `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB      `json:",omitempty,omitzero"`
}
