package member

type Repository interface {
	GeTProfileMemberRepo(UserID string) (*UserProfileModel, error)
	GetAddressByUserRepo(UserID string) (*[]AddressProfileModel, error)
	GetAddressByIdRepo(UserID string, AddressID int64) (*AddressProfileModel, error)
	AddProfileAddRepo(auth *AddAddressModel) (interface{}, error)
	UpdateProFileAddrRepo(auth *AddAddressModel) (interface{}, error)
	DeleteProfileAddressRepo(userID string, addressID int64) (interface{}, error)
	UpdatePRofileRepo(user *UpdateUserProfileModel, userID string) (interface{}, error)
	GetCouponUsers(userCode string) ([]ModelCoupon, error)
	CheckCouponRepo(userID, couponNO string) (interface{}, error)
	GetAllMember() (interface{}, error)
	RemoveUserByAdmin(adminID string, userID string) (interface{}, error)
	GetUserByAdmin(userID string) (*UserProfileModelBackEnd, error)
	UpdateMemberByadmin(user *UpdateUserProfileModelByAdmin, AdminID string) (interface{}, error)
	RemoveMember(adminID string, userID string, status int64) (interface{}, error)
}
