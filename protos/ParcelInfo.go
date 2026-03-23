package protos

type ParcelInfo struct {
	IParcelInfo
	Key         *ParcelKeyPair `json:",omitempty,omitzero"`
	Amount      int64          `json:",omitempty,omitzero"`
	Multiplier  *BasisPoint    `json:",omitempty,omitzero"`
	Probability *BasisPoint    `json:",omitempty,omitzero"`
}
