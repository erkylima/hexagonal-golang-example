package tests

import (
	"testing"

	"github.com/erkylima/hexagonal/hexagonal/internal/beneficiary"
	"github.com/erkylima/hexagonal/hexagonal/internal/beneficiary/serializer/json"

	"github.com/stretchr/testify/assert"
)

func TestBeneficiaryDecodeSerializer(t *testing.T) {
	expected := []byte(`{"name":"john","age":15, "address":"Rua", "phone":"879898"}`)
	beneficiaryExample := &beneficiary.Beneficiary{
		Name:    "john",
		Age:     15,
		Address: "Rua",
		Phone:   "879898",
	}
	b := &json.Beneficiary{}

	obteined, _ := b.Decode(expected)
	assert.Equal(t, beneficiaryExample, obteined)
}
