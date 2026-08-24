package protos

type TacticalRelayRankingInfoResponse struct {
	ResponsePacket
	RankingDB           *TacticalRelayRankingDB `json:",omitempty,omitzero"`
	IsRankStillCounting bool                    `json:",omitempty,omitzero"`
}
