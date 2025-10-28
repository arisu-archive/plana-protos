package protos

type RaidListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CreateRaidDBs []RaidDB `json:",omitempty,omitzero"`
	EnterRaidDBs []RaidDB `json:",omitempty,omitzero"`
	ListRaidDBs []RaidDB `json:",omitempty,omitzero"`
}
