package protos

type CafeProductionDB_CafeProductionParcelInfo struct {
	Key    ParcelKeyPair `json:",omitempty,omitzero"`
	Amount int64         `json:",omitempty,omitzero"`
}
