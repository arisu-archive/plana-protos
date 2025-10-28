package protos

type WeekDungeonBattleResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	WeekDungeonStageHistoryDB WeekDungeonStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs []CharacterDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
