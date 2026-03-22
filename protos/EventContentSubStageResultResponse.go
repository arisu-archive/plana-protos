package protos

type EventContentSubStageResultResponse struct {
	ResponsePacket
	TacticRank                int64 `json:",omitempty,omitzero"`
	CampaignStageHistoryDB    CampaignStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	ParcelResultDB            ParcelResultDB
	FirstClearReward          []ParcelInfo
	BonusReward               []ParcelInfo
	EventContentCollectionDBs []EventContentCollectionDB
}
