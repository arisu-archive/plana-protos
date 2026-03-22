package protos

type MiniGameDreamMakerDailyClosingResponse struct {
	ResponsePacket
	InfoDB                       MiniGameDreamMakerInfoDB
	ParameterDBs                 []MiniGameDreamMakerParameterDB
	ParcelResultDB               ParcelResultDB
	EventPointAmount             int64 `json:",omitempty,omitzero"`
	AlreadyReceivePointRewardIds []int64
}
