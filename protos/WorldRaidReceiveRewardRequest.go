package protos

type WorldRaidReceiveRewardRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
	PhaseId  int64 `json:",omitempty,omitzero"`
}
