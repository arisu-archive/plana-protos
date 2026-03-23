package protos

type ClanAssistRewardInfo struct {
	CharacterDBId            int64  `json:",omitempty,omitzero"`
	DeployDate               MxTime `json:",omitempty,omitzero"`
	RentCount                int64  `json:",omitempty,omitzero"`
	CumultativeRewardParcels []*ParcelInfo
	RentRewardParcels        []*ParcelInfo
}
