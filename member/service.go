package member

type service struct {
	repo Repository
}

// NewService creates new auth service
func NewService(member Repository) (Service, error) {
	s := service{member}
	return &s, nil
}

type Service interface {
	GetProfileMemberService(UserCode string) (interface{}, error)
	GetAddressByUserService(UserCode string) (interface{}, error)
	GetAddressByIdService(UserCode string, AddressID int64) (interface{}, error)
	AddProfileAddressService(auth *AddAddressModel) (interface{}, error)
	UpdateProfileAddress(auth *AddAddressModel) (interface{}, error)
	DeleteProfileAddressService(userID string, addressID int64) (interface{}, error)
	UpdatePRofileUserService(user *UpdateUserProfileModel, userID string) (interface{}, error)
	GetCouponService(userCode string) (interface{}, error)
	CheckCouponService(userID string, couponNO string) (interface{}, error)
	GetAllMember() (interface{}, error)
	RemoveUserServiceAdmin(AdminID string, userID string) (interface{}, error)
	GetUserByAdmin(userID string) (interface{}, error)
	UpdateMemberByAdmin(user *UpdateUserProfileModelByAdmin, AdminID string) (interface{}, error)
}

func (s *service) GetUserByAdmin(userID string) (interface{}, error) {
	resp, err := s.repo.GetUserByAdmin(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) UpdateMemberByAdmin(user *UpdateUserProfileModelByAdmin, AdminID string) (interface{}, error) {
	if user.ActiveStatus != 1 {
		resp, err := s.repo.RemoveMember(AdminID, user.UserID, user.ActiveStatus)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else {
		resp, err := s.repo.UpdateMemberByadmin(user, AdminID)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	return nil, nil
}

func (s *service) RemoveUserServiceAdmin(AdminID string, userID string) (interface{}, error) {
	resp, err := s.repo.RemoveUserByAdmin(AdminID, userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) CheckCouponService(userID string, couponNO string) (interface{}, error) {
	resp, err := s.repo.CheckCouponRepo(userID, couponNO)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) UpdatePRofileUserService(user *UpdateUserProfileModel, userID string) (interface{}, error) {
	_, err := s.repo.UpdatePRofileRepo(user, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) GetCouponService(userCode string) (interface{}, error) {
	profile, err := s.repo.GetCouponUsers(userCode)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *service) GetAddressByIdService(UserCode string, AddressID int64) (interface{}, error) {
	// ad := AddressProfileModel{
	// 	AddressID:   1,
	// 	Name:        "admin",
	// 	Phone:       "+661213131231",
	// 	Address:     "123/7 บ้านตำหนัก หมู่3 ต.ดอนแก้ว",
	// 	Province:    "chiang mai",
	// 	District:    "อ.สารภี",
	// 	PostalCode:  "50140",
	// 	MainAddress: 1,
	// }

	// return ad, nil
	profile, err := s.repo.GetAddressByIdRepo(UserCode, AddressID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *service) GetProfileMemberService(UserCode string) (interface{}, error) {
	profile, err := s.repo.GeTProfileMemberRepo(UserCode)
	if err != nil {
		return nil, err
	}
	// pr := ProfileMemberModel{
	// 	Email:     "admin@gmail.com",
	// 	Phone:     "+66121312345",
	// 	ShopName:  "admin",
	// 	Gender:    0,
	// 	BrithDate: "1994/04/15",
	// }
	// ad := []AddressProfileModel{
	// 	AddressProfileModel{
	// 		AddressID:   1,
	// 		Name:        "admin",
	// 		Phone:       "+661213131231",
	// 		Address:     "123/7 บ้านตำหนัก หมู่3 ต.ดอนแก้ว",
	// 		Province:    "chiang mai",
	// 		District:    "อ.สารภี",
	// 		PostalCode:  "50140",
	// 		MainAddress: 1,
	// 	},
	// 	AddressProfileModel{
	// 		AddressID:   2,
	// 		Name:        "admin",
	// 		Phone:       "+66123456789",
	// 		Address:     "123/7 บ้านตำหนัก หมู่3 ต.ดอนแก้ว",
	// 		Province:    "chiang mai",
	// 		District:    "อ.สารภี",
	// 		PostalCode:  "50140",
	// 		MainAddress: 0,
	// 	},
	// }

	// profile := ProfileModel{
	// 	UserID:  "1232131",
	// 	Profile: pr,
	// 	Address: ad,
	// }
	return profile, nil
}

func (s *service) GetAddressByUserService(UserCode string) (interface{}, error) {
	// ad := []AddressProfileModel{
	// 	AddressProfileModel{
	// 		AddressID:   1,
	// 		Name:        "admin",
	// 		Phone:       "+661213131231",
	// 		Address:     "123/7 บ้านตำหนัก หมู่3 ต.ดอนแก้ว",
	// 		Province:    "chiang mai",
	// 		District:    "อ.สารภี",
	// 		PostalCode:  "50140",
	// 		MainAddress: 1,
	// 	},
	// 	AddressProfileModel{
	// 		AddressID:   2,
	// 		Name:        "admin",
	// 		Phone:       "+66123456789",
	// 		Address:     "123/7 บ้านตำหนัก หมู่3 ต.ดอนแก้ว",
	// 		Province:    "chiang mai",
	// 		District:    "อ.สารภี",
	// 		PostalCode:  "50140",
	// 		MainAddress: 0,
	// 	},
	// }
	// return ad, nil
	profile, err := s.repo.GetAddressByUserRepo(UserCode)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *service) AddProfileAddressService(auth *AddAddressModel) (interface{}, error) {
	_, err := s.repo.AddProfileAddRepo(auth)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) UpdateProfileAddress(auth *AddAddressModel) (interface{}, error) {
	_, err := s.repo.UpdateProFileAddrRepo(auth)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) DeleteProfileAddressService(userID string, addressID int64) (interface{}, error) {
	_, err := s.repo.DeleteProfileAddressRepo(userID, addressID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) GetAllMember() (interface{}, error) {
	resp, err := s.repo.GetAllMember()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
