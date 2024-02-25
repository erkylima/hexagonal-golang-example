package beneficiary

import (
	"errors"

	"gopkg.in/dealancer/validate.v2"

	errs "github.com/pkg/errors"
)

var (
	ErrBeneficiaryNotFound = errors.New("Beneficiary Not Found")
	ErrBeneficiaryInvalid  = errors.New("Beneficiary Invalid")
)

type beneficiaryService struct {
	repository BeneficiaryRepository
}

func NewBeneficiaryService(repository BeneficiaryRepository) BeneficiaryService {
	return &beneficiaryService{
		repository,
	}
}

func (b *beneficiaryService) Find(name string) (*Beneficiary, error) {
	return b.repository.Find(name)
}

func (b *beneficiaryService) Store(beneficiary *Beneficiary) error {
	err := validate.Validate(beneficiary)
	if err != nil {
		return errs.Wrap(ErrBeneficiaryInvalid, "service.Beneficiary.Store")
	}
	return b.repository.Store(beneficiary)
}
