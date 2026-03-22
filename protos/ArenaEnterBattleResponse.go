package protos

type ArenaEnterBattleResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB
	ArenaBattleDB     ArenaBattleDB
	ArenaPlayerInfoDB ArenaPlayerInfoDB
	VictoryRewards    ParcelResultDB
	SeasonRewards     ParcelResultDB
	AllTimeRewards    ParcelResultDB
}
