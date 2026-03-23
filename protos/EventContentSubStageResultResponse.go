package protos

type EventContentSubStageResultResponse struct {
	ResponsePacket
	TacticRank                int64                   `json:",omitempty,omitzero"`
	CampaignStageHistoryDB    *CampaignStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs       []*CharacterDB
	ParcelResultDB            *ParcelResultDB `json:",omitempty,omitzero"`
	FirstClearReward          []*ParcelInfo
	BonusReward               []*ParcelInfo
	EventContentCollectionDBs []*EventContentCollectionDB
}
