package protos

type ShopGachaHistoryResponse struct {
	ResponsePacket
	GachaHistoryDBs []*GachaHistoryDB
}
