package protos

type MailReceiveSemiPermanentResponse struct {
	ResponsePacket
	MailDBId                           int64                        `json:",omitempty,omitzero"`
	ParcelResultDB                     *ParcelResultDB              `json:",omitempty,omitzero"`
	AppliedMonthlyProductPurchaseDB    *MonthlyProductPurchaseDB    `json:",omitempty,omitzero"`
	AppliedDailyRecordDB               *DailyRecordDB               `json:",omitempty,omitzero"`
	AppliedBattlePassProductPurchaseDB *BattlePassProductPurchaseDB `json:",omitempty,omitzero"`
	AppliedBattlePassInfoDB            *BattlePassInfoDB            `json:",omitempty,omitzero"`
	BattlePassInfoDBs                  []BattlePassInfoDB
}
