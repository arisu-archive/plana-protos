package protos

type ClanApplicantResponse struct {
	ResponsePacket
	ClanMemberDBs []*ClanMemberDB
}
