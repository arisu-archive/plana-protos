package protos

type HexaDisplayType int32

const (
	HexaDisplayType_None HexaDisplayType = 0
	HexaDisplayType_EndBattle HexaDisplayType = 1
	HexaDisplayType_PlayScenario HexaDisplayType = 2
	HexaDisplayType_SpawnUnitFromUniqueId HexaDisplayType = 3
	HexaDisplayType_StatBuff HexaDisplayType = 4
	HexaDisplayType_DieUnit HexaDisplayType = 5
	HexaDisplayType_HideStrategy HexaDisplayType = 6
	HexaDisplayType_SpawnUnit HexaDisplayType = 7
	HexaDisplayType_SpawnStrategy HexaDisplayType = 8
	HexaDisplayType_SpawnTile HexaDisplayType = 9
	HexaDisplayType_HideTile HexaDisplayType = 10
	HexaDisplayType_ClearFogOfWar HexaDisplayType = 11
	HexaDisplayType_MoveUnit HexaDisplayType = 12
	HexaDisplayType_WarpUnit HexaDisplayType = 13
	HexaDisplayType_SetTileMovablity HexaDisplayType = 14
	HexaDisplayType_WarpUnitFromHideTile HexaDisplayType = 15
	HexaDisplayType_BossExile HexaDisplayType = 16
)
