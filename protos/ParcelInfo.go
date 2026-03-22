package protos

type ParcelInfo struct {
	IParcelInfo
	Key         ParcelKeyPair
	Amount      int64 `json:",omitempty,omitzero"`
	Multiplier  BasisPoint
	Probability BasisPoint
}
