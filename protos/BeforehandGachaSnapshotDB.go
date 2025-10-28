package protos

type BeforehandGachaSnapshotDB struct {
	ShopUniqueId int64 `json:",omitempty,omitzero"`
	GoodsId int64 `json:",omitempty,omitzero"`
	LastIndex int64 `json:",omitempty,omitzero"`
	LastResults []int64 `json:",omitempty,omitzero"`
	SavedIndex *int64 `json:",omitempty,omitzero"`
	SavedResults []int64 `json:",omitempty,omitzero"`
	PickedIndex *int64 `json:",omitempty,omitzero"`
}
