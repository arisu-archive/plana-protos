package protos

type EventContentBonusRewardDB struct {
	EventContentId     int64      `json:",omitempty,omitzero"`
	EventStageUniqueId int64      `json:",omitempty,omitzero"`
	BonusPercentage    int64      `json:",omitempty,omitzero"`
	BonusParcelInfo    ParcelInfo `json:",omitempty,omitzero"`
}
