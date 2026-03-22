package protos

type MiniGameDreamMakerGetInfoResponse struct {
	ResponsePacket
	InfoDB                       MiniGameDreamMakerInfoDB
	ParameterDBs                 []MiniGameDreamMakerParameterDB
	EndingDBs                    []MiniGameDreamMakerEndingDB
	EventContentCollectionDBs    []EventContentCollectionDB
	EventPointAmount             int64 `json:",omitempty,omitzero"`
	AlreadyReceivePointRewardIds []int64
}
