package admin

type Repository interface {
	FindUserByEmail(code string) (*UserModelEmail, error)
	SaveToekenAdmin(email string, token string) (interface{}, error)
	GetProfileAdmin(admincode string) (*AdminProfileModel, error)
	RemoveToken(tokenID string) (interface{}, error)
	AddUserAdminRepo(adminID string, email string, pass string, roleID int64) (interface{}, error)
	UpdateProfileBackEndRepo(user *UpdateUserProfileModel, adminID string) (interface{}, error)
	FindAdminBack(email string) (bool, error)
	SigupUserBack(admin, email, pass string, roleID int64) (interface{}, error)
	ListAdminRepo(status int64, search string, limit int64) (interface{}, error)
	CountAdminRepo(status int64, search string) (int64, error)
	ListRoleRepo(status int64) (interface{}, error)
	UpdatePassword(user *UpdateUserProfileModel, pass string, adminID string) (interface{}, error)
	AddadminUser(adminID string, password string, user *AddadminBackEndModel) (interface{}, error)
	GetAdminByIDRepo(adminID string) (interface{}, error)
	RemoveAdmin(adminID string, status int64, adminCode string) (interface{}, error)
}
