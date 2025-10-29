package protos

type EliminateRaidGetBestTeamRequest struct {
	RequestPacket
	SearchAccountId int64 `json:",omitempty,omitzero"`
}
