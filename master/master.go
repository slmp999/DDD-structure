package master

//  BankModel  table model
type BankModel struct {
	ID       string `json:"user_id"`
	BankCode string `json:"member_profile"`
	BankName string `json:"member_address"`
}
