package protos

type MailReceiveSemiPermanentResponse struct {
	ResponsePacket
	MailDBId                           int64 `json:",omitempty,omitzero"`
	ParcelResultDB                     ParcelResultDB
	AppliedMonthlyProductPurchaseDB    MonthlyProductPurchaseDB
	AppliedDailyRecordDB               DailyRecordDB
	AppliedBattlePassProductPurchaseDB BattlePassProductPurchaseDB
	AppliedBattlePassInfoDB            BattlePassInfoDB
	BattlePassInfoDBs                  []BattlePassInfoDB
}
