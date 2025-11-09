package protos

type ConquestTriggerType int32

const (
	ConquestTriggerType_None                   ConquestTriggerType = 0
	ConquestTriggerType_TileConquer            ConquestTriggerType = 1
	ConquestTriggerType_TileUpgrade            ConquestTriggerType = 2
	ConquestTriggerType_MapEnter               ConquestTriggerType = 3
	ConquestTriggerType_SyncState              ConquestTriggerType = 4
	ConquestTriggerType_AcquireCalculateReward ConquestTriggerType = 5
	ConquestTriggerType_UnexpectedEvent        ConquestTriggerType = 6
	ConquestTriggerType_MassErosion            ConquestTriggerType = 7
	ConquestTriggerType_MassErosionEnd         ConquestTriggerType = 8
	ConquestTriggerType_TileErosion            ConquestTriggerType = 9
	ConquestTriggerType_TileErosionEnd         ConquestTriggerType = 10
)
