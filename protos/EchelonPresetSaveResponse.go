package protos

type EchelonPresetSaveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetDB EchelonPresetDB `json:",omitempty,omitzero"`
}
