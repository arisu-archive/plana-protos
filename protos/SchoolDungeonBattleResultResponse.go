package protos

type SchoolDungeonBattleResultResponse struct {
	ResponsePacket
	SchoolDungeonStageHistoryDB SchoolDungeonStageHistoryDB `json:",omitempty,omitzero"`
	LevelUpCharacterDBs         []CharacterDB               `json:",omitempty,omitzero"`
	FirstClearReward            []ParcelInfo                `json:",omitempty,omitzero"`
	ThreeStarReward             []ParcelInfo                `json:",omitempty,omitzero"`
	ParcelResultDB              ParcelResultDB              `json:",omitempty,omitzero"`
}
