package protos

type WeekDungeonBattleResultResponse struct {
	ResponsePacket
	WeekDungeonStageHistoryDB *WeekDungeonStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs       []CharacterDB
	ParcelResultDB            *ParcelResultDB `json:",omitempty,omitzero"`
}
