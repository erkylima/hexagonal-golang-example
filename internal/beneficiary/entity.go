package beneficiary

type Beneficiary struct {
	Name    string `json:"name" bson:"name" validate:"empty=false"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
