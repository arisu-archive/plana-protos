package protos

type StickerDB struct {
	ParcelBase
	StickerUniqueId int64 `json:",omitempty,omitzero"`
}
