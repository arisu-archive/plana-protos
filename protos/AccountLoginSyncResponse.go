package protos

type AccountLoginSyncResponse struct {
	ResponsePacket
	CafeGetInfoResponse                      *CafeGetInfoResponse                      `json:",omitempty,omitzero"`
	AccountCurrencySyncResponse              *AccountCurrencySyncResponse              `json:",omitempty,omitzero"`
	CharacterListResponse                    *CharacterListResponse                    `json:",omitempty,omitzero"`
	EquipmentItemListResponse                *EquipmentItemListResponse                `json:",omitempty,omitzero"`
	CharacterGearListResponse                *CharacterGearListResponse                `json:",omitempty,omitzero"`
	ItemListResponse                         *ItemListResponse                         `json:",omitempty,omitzero"`
	EchelonListResponse                      *EchelonListResponse                      `json:",omitempty,omitzero"`
	MemoryLobbyListResponse                  *MemoryLobbyListResponse                  `json:",omitempty,omitzero"`
	CampaignListResponse                     *CampaignListResponse                     `json:",omitempty,omitzero"`
	ArenaLoginResponse                       *ArenaLoginResponse                       `json:",omitempty,omitzero"`
	RaidLoginResponse                        *RaidLoginResponse                        `json:",omitempty,omitzero"`
	EliminateRaidLoginResponse               *EliminateRaidLoginResponse               `json:",omitempty,omitzero"`
	CraftInfoListResponse                    *CraftInfoListResponse                    `json:",omitempty,omitzero"`
	ClanLoginResponse                        *ClanLoginResponse                        `json:",omitempty,omitzero"`
	MomotalkOutlineResponse                  *MomoTalkOutLineResponse                  `json:",omitempty,omitzero"`
	ScenarioListResponse                     *ScenarioListResponse                     `json:",omitempty,omitzero"`
	ShopGachaRecruitListResponse             *ShopGachaRecruitListResponse             `json:",omitempty,omitzero"`
	TimeAttackDungeonLoginResponse           *TimeAttackDungeonLoginResponse           `json:",omitempty,omitzero"`
	BillingPurchaseListByYostarResponse      *BillingPurchaseListByYostarResponse      `json:",omitempty,omitzero"`
	EventContentPermanentListResponse        *EventContentPermanentListResponse        `json:",omitempty,omitzero"`
	AttachmentGetResponse                    *AttachmentGetResponse                    `json:",omitempty,omitzero"`
	AttachmentEmblemListResponse             *AttachmentEmblemListResponse             `json:",omitempty,omitzero"`
	ContentSweepMultiSweepPresetListResponse *ContentSweepMultiSweepPresetListResponse `json:",omitempty,omitzero"`
	StickerListResponse                      *StickerLoginResponse                     `json:",omitempty,omitzero"`
	MultiFloorRaidSyncResponse               *MultiFloorRaidSyncResponse               `json:",omitempty,omitzero"`
	MultiFloorRaidLoginResponse              *MultiFloorRaidLoginResponse              `json:",omitempty,omitzero"`
	FriendCount                              int64                                     `json:",omitempty,omitzero"`
	FriendCode                               string
	PickupFirstGetHistoryDBs                 []*PickupFirstGetHistoryDB
	AccountLevelRewardIds                    []int64
}
