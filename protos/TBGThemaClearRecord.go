package protos

type TBGThemaClearRecord struct {
	ThemaIndex int32        `json:",omitempty,omitzero"`
	SweepCosts []ParcelInfo `json:",omitempty,omitzero"`
}
