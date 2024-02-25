package beneficiary

type BeneficiaryRepository interface {
	Find(name string) (*Beneficiary, error)
	Store(*Beneficiary) error
}
