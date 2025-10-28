package protos

type CharacterGearListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GearDBs []GearDB `json:",omitempty,omitzero"`
}
