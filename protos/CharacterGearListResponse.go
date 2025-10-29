package protos

type CharacterGearListResponse struct {
	ResponsePacket
	GearDBs []GearDB `json:",omitempty,omitzero"`
}
