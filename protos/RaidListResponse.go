package protos

type RaidListResponse struct {
	ResponsePacket
	CreateRaidDBs []*RaidDB
	EnterRaidDBs  []*RaidDB
	ListRaidDBs   []*RaidDB
}
