package protos

type TacticalRelayEndBattleResponse struct {
	ResponsePacket
	ParcelResultDB  *ParcelResultDB `json:",omitempty,omitzero"`
	ClearHistoryDBs []*TacticalRelayHistoryDB
	IsNewRecord     bool `json:",omitempty,omitzero"`
}
