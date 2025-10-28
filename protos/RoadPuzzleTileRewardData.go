package protos

type RoadPuzzleTileRewardData struct {
	Location HexLocation `json:",omitempty,omitzero"`
	RewardInfo ParcelInfo `json:",omitempty,omitzero"`
}
