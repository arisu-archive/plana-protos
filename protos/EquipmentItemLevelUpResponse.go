package protos

type EquipmentItemLevelUpResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EquipmentDB EquipmentDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
