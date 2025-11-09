package protos

type MiniGameDreamMakerDailyClosingResponse struct {
	ResponsePacket
	InfoDB                       MiniGameDreamMakerInfoDB        `json:",omitempty,omitzero"`
	ParameterDBs                 []MiniGameDreamMakerParameterDB `json:",omitempty,omitzero"`
	ParcelResultDB               ParcelResultDB                  `json:",omitempty,omitzero"`
	EventPointAmount             int64                           `json:",omitempty,omitzero"`
	AlreadyReceivePointRewardIds []int64                         `json:",omitempty,omitzero"`
}
