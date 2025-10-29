package protos

type EliminateRaidSeasonRewardResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ReceiveRewardIds []int64 `json:",omitempty,omitzero"`
}
