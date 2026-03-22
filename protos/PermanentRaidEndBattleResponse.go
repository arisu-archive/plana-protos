package protos

type PermanentRaidEndBattleResponse struct {
	ResponsePacket
	ScoreInfo       RaidScoreInfo
	BattleHistoryDB PermanentRaidBattleHistoryDB
}
