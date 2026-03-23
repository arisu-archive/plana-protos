package protos

type BillingPurchaseListByYostarResponse struct {
	ResponsePacket
	CountList                 []*PurchaseCountDB
	OrderList                 []*PurchaseOrderDB
	MonthlyProductList        []*MonthlyProductPurchaseDB
	BlockedProductDBs         []*BlockedProductDB
	BattlePassProductList     []*BattlePassProductPurchaseDB
	PendingProductShopCashIds []int64
}
