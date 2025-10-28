package protos

type CharacterGearTierUpRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GearServerId int64 `json:",omitempty,omitzero"`
	ReplaceInfos []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
