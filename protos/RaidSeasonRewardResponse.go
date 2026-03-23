package protos

type RaidSeasonRewardResponse struct {
	ResponsePacket
	ParcelResultDB   *ParcelResultDB `json:",omitempty,omitzero"`
	ReceiveRewardIds []int64
}
