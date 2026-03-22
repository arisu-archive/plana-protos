package protos

type CafeProductionDB_CafeProductionParcelInfo struct {
	Key    ParcelKeyPair
	Amount int64 `json:",omitempty,omitzero"`
}
