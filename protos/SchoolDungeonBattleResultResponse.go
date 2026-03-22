package protos

type SchoolDungeonBattleResultResponse struct {
	ResponsePacket
	SchoolDungeonStageHistoryDB SchoolDungeonStageHistoryDB
	LevelUpCharacterDBs         []CharacterDB
	FirstClearReward            []ParcelInfo
	ThreeStarReward             []ParcelInfo
	ParcelResultDB              ParcelResultDB
}
