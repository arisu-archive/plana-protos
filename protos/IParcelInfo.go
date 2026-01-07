package protos

type IParcelInfo struct {
	Key    ParcelKeyPair `json:",omitempty,omitzero"`
	Amount int64         `json:",omitempty,omitzero"`
}
