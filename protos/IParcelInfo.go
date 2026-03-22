package protos

type IParcelInfo struct {
	Key    ParcelKeyPair
	Amount int64 `json:",omitempty,omitzero"`
}
