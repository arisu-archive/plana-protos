package protos

type CharacterGearTierUpRequest struct {
	RequestPacket
	GearServerId int64                     `json:",omitempty,omitzero"`
	ReplaceInfos []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
