package protos

type ClanCreateResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
}
