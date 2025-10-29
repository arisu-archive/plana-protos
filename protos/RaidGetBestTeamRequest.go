package protos

type RaidGetBestTeamRequest struct {
	RequestPacket
	SearchAccountId int64 `json:",omitempty,omitzero"`
}
