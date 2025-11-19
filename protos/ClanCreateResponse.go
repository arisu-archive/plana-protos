package protos

type ClanCreateResponse struct {
	ResponsePacket
	ClanDB            ClanDB            `json:",omitempty,omitzero"`
	ClanMemberDB      ClanMemberDB      `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
}
