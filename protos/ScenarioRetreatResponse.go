package protos

type ScenarioRetreatResponse struct {
	ResponsePacket
	ReleasedEchelonNumbers []int64
	ParcelResultDB         ParcelResultDB
}
