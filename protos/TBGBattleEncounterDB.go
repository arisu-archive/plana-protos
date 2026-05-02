package protos

type TBGBattleEncounterDB struct {
	TBGEncounterDB
	Stage TBGBattleEncounterDB_TBGBattleEncounterStage `json:"st,omitempty,omitzero"`
}
