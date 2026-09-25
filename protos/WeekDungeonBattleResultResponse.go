package protos

type WeekDungeonBattleResultResponse struct {
	ResponsePacket
	WeekDungeonStageHistoryDB *WeekDungeonStageHistoryDB `json:",omitempty,omitzero"`
	PrevStageHistoryDBs       []*WeekDungeonStageHistoryDB
	LevelUpCharacterDBs       []*CharacterDB
	ParcelResultDB            *ParcelResultDB `json:",omitempty,omitzero"`
}
