package protos

type MiniGameTankRaceLevelUpResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB                    `json:",omitempty,omitzero"`
	TankDB         *MiniGameTankRaceTankUpgradeInfoDB `json:",omitempty,omitzero"`
}
