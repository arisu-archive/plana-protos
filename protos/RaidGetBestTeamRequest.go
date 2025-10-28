package protos

type RaidGetBestTeamRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SearchAccountId int64 `json:",omitempty,omitzero"`
}
