package protos

type WeekDungeonListResponse struct {
	ResponsePacket
	AdditionalStageIdList         []int64
	WeekDungeonStageHistoryDBList []*WeekDungeonStageHistoryDB
}
