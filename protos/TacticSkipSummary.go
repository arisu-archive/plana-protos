package protos

type TacticSkipSummary struct {
	StageId           int64 `json:",omitempty,omitzero"`
	Group01HexaUnitId int64 `json:",omitempty,omitzero"`
	Group02HexaUnitId int64 `json:",omitempty,omitzero"`
}
