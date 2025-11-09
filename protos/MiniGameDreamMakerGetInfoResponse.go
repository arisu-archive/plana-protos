package protos

type MiniGameDreamMakerGetInfoResponse struct {
	ResponsePacket
	InfoDB                       MiniGameDreamMakerInfoDB        `json:",omitempty,omitzero"`
	ParameterDBs                 []MiniGameDreamMakerParameterDB `json:",omitempty,omitzero"`
	EndingDBs                    []MiniGameDreamMakerEndingDB    `json:",omitempty,omitzero"`
	EventContentCollectionDBs    []EventContentCollectionDB      `json:",omitempty,omitzero"`
	EventPointAmount             int64                           `json:",omitempty,omitzero"`
	AlreadyReceivePointRewardIds []int64                         `json:",omitempty,omitzero"`
}
