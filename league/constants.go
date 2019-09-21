package league

type Queue string

const (
	RankedSolo5x5 Queue = "RANKED_SOLO_5x5"
	RankedTFT     Queue = "RANKED_TFT"
	RankedFlexSR  Queue = "RANKED_FLEX_SR"
	RankedFlexTT  Queue = "RANKED_FLEX_TT"
)

func (q Queue) String() string {
	return string(q)
}

type TierExp string

const (
	ExpChallenger  TierExp = "CHALLENGER"
	ExpGrandMaster TierExp = "GRADMASTER"
	ExpMaster      TierExp = "MASTER"
	ExpDiamond     TierExp = "DIAMOND"
	ExpPlatinum    TierExp = "PLATINUM"
	ExpGold        TierExp = "GOLD"
	ExpSilver      TierExp = "SILVER"
	ExpBronze      TierExp = "BRONZE"
	ExpIron        TierExp = "IRON"
)

func (te TierExp) String() string {
	return string(te)
}

type Tier string

const (
	Diamond  Tier = "DIAMOND"
	Platinum Tier = "PLATINUM"
	Gold     Tier = "GOLD"
	Silver   Tier = "SILVER"
	Bronze   Tier = "BRONZE"
	Iron     Tier = "IRON"
)

func (t Tier) String() string {
	return string(t)
}

type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
)

func (d Division) String() string {
	return string(d)
}
