package protos

type WeekDungeonListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AdditionalStageIdList []int64 `json:",omitempty,omitzero"`
	WeekDungeonStageHistoryDBList []WeekDungeonStageHistoryDB `json:",omitempty,omitzero"`
}
