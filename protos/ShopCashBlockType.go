package protos

type ShopCashBlockType int32

const (
	ShopCashBlockType_None       ShopCashBlockType = -9999
	ShopCashBlockType_GooglePlay ShopCashBlockType = -3
	ShopCashBlockType_AppStore   ShopCashBlockType = -2
	ShopCashBlockType_All        ShopCashBlockType = -1
)
