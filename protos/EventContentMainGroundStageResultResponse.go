package protos

type EventContentMainGroundStageResultResponse struct {
	ResponsePacket
	TacticRank                int64 `json:",omitempty,omitzero"`
	CampaignStageHistoryDB    CampaignStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	ParcelResultDB            ParcelResultDB
	FirstClearReward          []ParcelInfo
	ThreeStarReward           []ParcelInfo
	BonusReward               []ParcelInfo
	EventContentCollectionDBs []EventContentCollectionDB
}
