package protos

type ShopEligmaHistoryDB struct {
	CharacterUniqueId int64 `json:",omitempty,omitzero"`
	PurchaseCount     int64 `json:",omitempty,omitzero"`
}
