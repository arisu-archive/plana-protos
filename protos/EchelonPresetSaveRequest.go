package protos

type EchelonPresetSaveRequest struct {
	RequestPacket
	PresetDB EchelonPresetDB `json:",omitempty,omitzero"`
}
