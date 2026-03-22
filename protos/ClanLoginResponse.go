package protos

type ClanLoginResponse struct {
	ResponsePacket
	AccountClanDB       ClanDB
	AccountClanMemberDB ClanMemberDB
	ClanAssistSlotDBs   []ClanAssistSlotDB
}
