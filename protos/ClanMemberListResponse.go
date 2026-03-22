package protos

type ClanMemberListResponse struct {
	ResponsePacket
	ClanDB        ClanDB
	ClanMemberDBs []ClanMemberDB
}
