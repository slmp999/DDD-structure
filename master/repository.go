package master
//  define repository
type Repository interface {
	GetBankList() (interface{}, error)
}
