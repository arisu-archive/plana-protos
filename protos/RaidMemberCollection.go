package protos

type RaidMemberCollectionItem struct {
	RaidMemberDescription
	TotalDamage int64
	RaidDamages []RaidDamage
}

type RaidMemberCollection []RaidMemberCollectionItem
