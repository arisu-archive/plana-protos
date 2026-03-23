package protos

type ArenaHistoryResponse struct {
	ResponsePacket
	ArenaHistoryDBs     []*ArenaHistoryDB
	ArenaDamageReportDB []*ArenaDamageReportDB
}
