package protos

type ClanMemberResponse struct {
	ResponsePacket
	ClanDB                ClanDB
	ClanMemberDB          ClanMemberDB
	DetailedAccountInfoDB DetailedAccountInfoDB
}
