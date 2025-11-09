package protos

type WeekDungeonBattleResultResponse struct {
	ResponsePacket
	WeekDungeonStageHistoryDB WeekDungeonStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs       []CharacterDB             `json:",omitempty,omitzero"`
	ParcelResultDB            ParcelResultDB            `json:",omitempty,omitzero"`
}
