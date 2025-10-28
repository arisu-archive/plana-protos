package protos

type EchelonPresetSaveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetDB EchelonPresetDB `json:",omitempty,omitzero"`
}
