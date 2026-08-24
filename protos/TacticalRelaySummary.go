package protos

type TacticalRelaySummary struct {
	ClearedWaveCount int32 `json:",omitempty,omitzero"`
	SquadRecords     []*RelaySquadRecord
}
