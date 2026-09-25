package protos

type ClanApplicantResponse struct {
	ResponsePacket
	ClanMemberDBs         []*ClanMemberDB
	KickHistoryAccountIds []int64
}
