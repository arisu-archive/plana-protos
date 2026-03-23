package protos

type GuideMissionSeasonListResponse struct {
	ResponsePacket
	GuideMissionSeasonDBs []*GuideMissionSeasonDB
}
