package protos

type ArenaEnterBattleResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ArenaBattleDB     ArenaBattleDB     `json:",omitempty,omitzero"`
	ArenaPlayerInfoDB ArenaPlayerInfoDB `json:",omitempty,omitzero"`
	VictoryRewards    ParcelResultDB    `json:",omitempty,omitzero"`
	SeasonRewards     ParcelResultDB    `json:",omitempty,omitzero"`
	AllTimeRewards    ParcelResultDB    `json:",omitempty,omitzero"`
}
