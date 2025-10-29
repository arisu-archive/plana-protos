package protos

type ArenaRankListResponse struct {
	ResponsePacket
	TopRankedUserDBs []ArenaUserDB `json:",omitempty,omitzero"`
}
