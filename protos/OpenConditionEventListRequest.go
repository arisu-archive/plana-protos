package protos

type OpenConditionEventListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConquestEventIds []int64 `json:",omitempty,omitzero"`
	WorldRaidSeasonAndGroupIds map[int64]int64 `json:",omitempty,omitzero"`
}
