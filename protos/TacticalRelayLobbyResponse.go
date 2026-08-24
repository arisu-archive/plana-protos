package protos

type TacticalRelayLobbyResponse struct {
	ResponsePacket
	ClearHistoryDBs           []*TacticalRelayHistoryDB
	PrevSeasonClearHistoryDBs []*TacticalRelayHistoryDB
}
