package protos

type WeekDungeonBattleResultResponse struct {
	ResponsePacket
	WeekDungeonStageHistoryDB WeekDungeonStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	ParcelResultDB            ParcelResultDB
}
