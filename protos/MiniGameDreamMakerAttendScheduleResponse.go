package protos

type MiniGameDreamMakerAttendScheduleResponse struct {
	ResponsePacket
	InfoDB                    MiniGameDreamMakerInfoDB
	ParameterDBs              []MiniGameDreamMakerParameterDB
	ParcelResultDB            ParcelResultDB
	ScheduleResultId          int64 `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB
}
