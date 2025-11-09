package protos

type RaidListResponse struct {
	ResponsePacket
	CreateRaidDBs []RaidDB `json:",omitempty,omitzero"`
	EnterRaidDBs  []RaidDB `json:",omitempty,omitzero"`
	ListRaidDBs   []RaidDB `json:",omitempty,omitzero"`
}
