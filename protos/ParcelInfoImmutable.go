package protos

type ParcelInfoImmutable struct {
	IParcelInfo
	Key    ParcelKeyPair
	Amount int64 `json:",omitempty,omitzero"`
}
