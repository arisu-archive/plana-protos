package protos

type ClanPermitResponse struct {
	ResponsePacket
	ClanDB       ClanDB
	ClanMemberDB ClanMemberDB
}
