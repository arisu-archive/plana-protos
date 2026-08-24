package protos

type MultiSweepPresetDB struct {
	PresetId   int64 `json:",omitempty,omitzero"`
	PresetName string
	StageIds   []int64
	ParcelIds  []*ParcelKeyPair
}
