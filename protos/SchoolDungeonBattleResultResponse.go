package protos

type SchoolDungeonBattleResultResponse struct {
	ResponsePacket
	SchoolDungeonStageHistoryDB *SchoolDungeonStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs         []*CharacterDB
	FirstClearReward            []*ParcelInfo
	ThreeStarReward             []*ParcelInfo
	ParcelResultDB              *ParcelResultDB `json:",omitempty,omitzero"`
}
