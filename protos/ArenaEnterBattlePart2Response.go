package protos

type ArenaEnterBattlePart2Response struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ArenaBattleDB ArenaBattleDB `json:",omitempty,omitzero"`
	ArenaPlayerInfoDB ArenaPlayerInfoDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	VictoryRewards ParcelResultDB `json:",omitempty,omitzero"`
	SeasonRewards ParcelResultDB `json:",omitempty,omitzero"`
	AllTimeRewards ParcelResultDB `json:",omitempty,omitzero"`
}
