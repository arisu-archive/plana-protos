package protos

type EchelonPresetSaveResponse struct {
	ResponsePacket
	PresetDB EchelonPresetDB `json:",omitempty,omitzero"`
}
