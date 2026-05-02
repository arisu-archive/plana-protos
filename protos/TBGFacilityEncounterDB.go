package protos

type TBGFacilityEncounterDB struct {
	TBGEncounterDB
	Stage                 TBGFacilityEncounterDB_TBGFacilityEncounterStage `json:"st,omitempty,omitzero"`
	EncounterOptionChoice int32                                            `json:"ec,omitempty,omitzero"`
}
