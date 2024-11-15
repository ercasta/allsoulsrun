package engine

type Equipment struct {
	Name                     string
	PhysicalDamage           IntRange
	IceDamage                IntRange
	FireDamage               IntRange
	LightningDamage          IntRange
	DexterityBonus           IntRange
	StrengthBonus            IntRange
	IntelligenceBonus        IntRange
	ConstitutionBonus        IntRange
	DexterityBonusPercent    int
	StrengthBonusPercent     int
	IntelligenceBonusPercent int
	ConstitutionBonusPercent int
}
