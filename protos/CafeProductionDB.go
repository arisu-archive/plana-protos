package protos

type CafeProductionDB struct {
	CafeDBId              int64                                       `json:",omitempty,omitzero"`
	ComfortValue          int64                                       `json:",omitempty,omitzero"`
	AppliedDate           MxTime                                      `json:",omitempty,omitzero"`
	ProductionParcelInfos []CafeProductionDB_CafeProductionParcelInfo `json:",omitempty,omitzero"`
}
