package json

import (
	"encoding/json"

	"github.com/erkylima/hexagonal/hexagonal/internal/beneficiary"
	"github.com/pkg/errors"
)

type Beneficiary struct{}

func (b *Beneficiary) Decode(input []byte) (*beneficiary.Beneficiary, error) {
	beneficiary := &beneficiary.Beneficiary{}

	if err := json.Unmarshal(input, beneficiary); err != nil {
		return nil, errors.Wrap(err, "serializer.Beneficiary.Decode")
	}
	return beneficiary, nil

}

func (b *Beneficiary) Encode(input *beneficiary.Beneficiary) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Beneficiary.Encode")
	}
	return rawMsg, nil
}
