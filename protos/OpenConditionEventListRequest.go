package protos

type OpenConditionEventListRequest struct {
	RequestPacket
	ConquestEventIds           []int64         `json:",omitempty,omitzero"`
	WorldRaidSeasonAndGroupIds map[int64]int64 `json:",omitempty,omitzero"`
}
