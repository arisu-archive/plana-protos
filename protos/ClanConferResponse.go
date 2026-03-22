package protos

type ClanConferResponse struct {
	ResponsePacket
	ClanMemberDB            ClanMemberDB
	AccountClanMemberDB     ClanMemberDB
	ClanDB                  ClanDB
	ClanMemberDescriptionDB ClanMemberDescriptionDB
}
