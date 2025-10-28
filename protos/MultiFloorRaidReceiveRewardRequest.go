package protos

type MultiFloorRaidReceiveRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SeasonId int64 `json:",omitempty,omitzero"`
	RewardDifficulty int32 `json:",omitempty,omitzero"`
}
