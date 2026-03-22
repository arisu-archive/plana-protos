package protos

type ClanCreateResponse struct {
	ResponsePacket
	ClanDB            ClanDB
	ClanMemberDB      ClanMemberDB
	AccountCurrencyDB AccountCurrencyDB
}
