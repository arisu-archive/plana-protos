package protos

type OpenConditionLockReason int32

const (
	OpenConditionLockReason_None OpenConditionLockReason = 0
	OpenConditionLockReason_Level OpenConditionLockReason = 1
	OpenConditionLockReason_StageClear OpenConditionLockReason = 2
	OpenConditionLockReason_Time OpenConditionLockReason = 4
	OpenConditionLockReason_Day OpenConditionLockReason = 8
	OpenConditionLockReason_CafeRank OpenConditionLockReason = 16
	OpenConditionLockReason_ScenarioModeClear OpenConditionLockReason = 32
	OpenConditionLockReason_CafeOpen OpenConditionLockReason = 64
)
