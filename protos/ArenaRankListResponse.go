package protos

type ArenaRankListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TopRankedUserDBs []ArenaUserDB `json:",omitempty,omitzero"`
}
