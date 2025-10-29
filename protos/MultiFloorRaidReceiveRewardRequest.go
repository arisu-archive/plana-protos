package protos

type MultiFloorRaidReceiveRewardRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
	RewardDifficulty int32 `json:",omitempty,omitzero"`
}
