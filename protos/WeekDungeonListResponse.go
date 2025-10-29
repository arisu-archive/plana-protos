package protos

type WeekDungeonListResponse struct {
	ResponsePacket
	AdditionalStageIdList []int64 `json:",omitempty,omitzero"`
	WeekDungeonStageHistoryDBList []WeekDungeonStageHistoryDB `json:",omitempty,omitzero"`
}
