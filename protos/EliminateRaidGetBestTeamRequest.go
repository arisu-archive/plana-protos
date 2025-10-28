package protos

type EliminateRaidGetBestTeamRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SearchAccountId int64 `json:",omitempty,omitzero"`
}
