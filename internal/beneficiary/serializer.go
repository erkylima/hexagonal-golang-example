package beneficiary

type BeneficiarySerializer interface {
	Decode(input []byte) (*Beneficiary, error)
	Encode(input *Beneficiary) ([]byte, error)
}
