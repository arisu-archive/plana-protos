package protos

type MiniGameTankRaceLobbyResponse struct {
	ResponsePacket
	StageHistoryDBs []*MiniGameTankRaceHistoryDB
	TankDBs         []*MiniGameTankRaceTankUpgradeInfoDB
}
