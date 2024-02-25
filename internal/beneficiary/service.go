package beneficiary

type BeneficiaryService interface {
	Find(name string) (*Beneficiary, error)
	Store(beneficiary *Beneficiary) error
}
