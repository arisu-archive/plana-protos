package protos

type MiniGameDreamMakerGetInfoResponse struct {
	ResponsePacket
	InfoDB                       *MiniGameDreamMakerInfoDB `json:",omitempty,omitzero"`
	ParameterDBs                 []*MiniGameDreamMakerParameterDB
	EndingDBs                    []*MiniGameDreamMakerEndingDB
	EventContentCollectionDBs    []*EventContentCollectionDB
	EventPointAmount             int64 `json:",omitempty,omitzero"`
	AlreadyReceivePointRewardIds []int64
}
