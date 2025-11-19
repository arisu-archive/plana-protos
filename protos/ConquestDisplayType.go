package protos

type ConquestDisplayType int32

const (
	ConquestDisplayType_None                     ConquestDisplayType = 0
	ConquestDisplayType_TileConquered            ConquestDisplayType = 1
	ConquestDisplayType_TileUpgraded             ConquestDisplayType = 2
	ConquestDisplayType_UnexpectedEvent          ConquestDisplayType = 3
	ConquestDisplayType_BossOpen                 ConquestDisplayType = 4
	ConquestDisplayType_PropAnimation            ConquestDisplayType = 5
	ConquestDisplayType_PropAnimationAndBlock    ConquestDisplayType = 6
	ConquestDisplayType_PropAnimationHoldAndPlay ConquestDisplayType = 7
	ConquestDisplayType_Operator                 ConquestDisplayType = 8
	ConquestDisplayType_StepComplete             ConquestDisplayType = 9
	ConquestDisplayType_MassErosion              ConquestDisplayType = 10
	ConquestDisplayType_Erosion                  ConquestDisplayType = 11
	ConquestDisplayType_ErosionRemove            ConquestDisplayType = 12
	ConquestDisplayType_CheckTileErosion         ConquestDisplayType = 13
	ConquestDisplayType_StepOpen                 ConquestDisplayType = 14
	ConquestDisplayType_BossClear                ConquestDisplayType = 15
	ConquestDisplayType_HideConquestUI           ConquestDisplayType = 16
	ConquestDisplayType_ShowConquestUI           ConquestDisplayType = 17
	ConquestDisplayType_HideHexaUI               ConquestDisplayType = 18
	ConquestDisplayType_ShowHexaUI               ConquestDisplayType = 19
	ConquestDisplayType_StepObjectComplete       ConquestDisplayType = 20
	ConquestDisplayType_CameraSetting            ConquestDisplayType = 21
	ConquestDisplayType_PlayMapEnterScenario     ConquestDisplayType = 22
	ConquestDisplayType_ShowTileConquerReward    ConquestDisplayType = 23
)
