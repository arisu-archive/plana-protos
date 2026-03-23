package protos

type ParcelInfoImmutable struct {
	IParcelInfo
	Key    *ParcelKeyPair `json:",omitempty,omitzero"`
	Amount int64          `json:",omitempty,omitzero"`
}
