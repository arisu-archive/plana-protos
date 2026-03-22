package protos

type EquipmentItemLevelUpResponse struct {
	ResponsePacket
	EquipmentDB       EquipmentDB
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
}
