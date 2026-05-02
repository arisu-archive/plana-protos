package protos

type TBGRandomEncounterDB struct {
	TBGEncounterDB
	Stage                 TBGRandomEncounterDB_TBGRandomEncounterStage `json:"st,omitempty,omitzero"`
	EncounterOptionChoice int32                                        `json:"ec,omitempty,omitzero"`
}
