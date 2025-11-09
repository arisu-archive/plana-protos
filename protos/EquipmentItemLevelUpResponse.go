package protos

type EquipmentItemLevelUpResponse struct {
	ResponsePacket
	EquipmentDB       EquipmentDB       `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ConsumeResultDB   ConsumeResultDB   `json:",omitempty,omitzero"`
}
