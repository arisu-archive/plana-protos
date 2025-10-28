package protos

type EliminateRaidSeasonRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
