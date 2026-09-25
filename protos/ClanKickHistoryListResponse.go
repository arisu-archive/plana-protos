package protos

type ClanKickHistoryListResponse struct {
	ResponsePacket
	ClanKickHistoryDetailDBs []*ClanKickHistoryDetailDB
}
