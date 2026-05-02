package protos

type TBGTreasureBoxEncounterDB struct {
	TBGEncounterDB
	Stage          TBGTreasureBoxEncounterDB_TBGTreasureEncounterStage `json:"st,omitempty,omitzero"`
	IsRealTreasure bool                                                `json:"real,omitempty,omitzero"`
}
