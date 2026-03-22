package protos

type CraftPresetSlotDB struct {
	PresetIndex   int32 `json:",omitempty,omitzero"`
	PresetNodeDBs []CraftPresetNodeDB
	PresetName    string
}
