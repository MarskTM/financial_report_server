package dao

import (
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database/do"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDAO struct {
	db *gorm.DB
}

func (u *UserDAO) CreateUser(user do.User, profile do.Profile) (*do.User, error) {
	var userInfo do.User

	if err := u.db.Debug().Transaction(func(tx *gorm.DB) error {
		if err := u.db.Model(&user).Clauses(clause.Returning{}).
			Create(&user).Error; err != nil {
			return err
		}

		if err := u.db.Model(&do.UserRole{}).Create(&do.UserRole{
			UserID: user.ID,
			RoleID: int32(env.ClientType),
			Active: true,
		}).Error; err != nil {
			return err
		}

		if err := u.db.Model(&do.Profile{}).Create(&do.Profile{
			UserID:    user.ID,
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
			Email:     profile.Email,
			Phone:     profile.Phone,
			Birthdate: profile.Birthdate,
		}).Error; err != nil {
			return err
		}

		if err := u.db.Model(&do.User{}).Where("id = ?", user.ID).Preload("UserRoles.Role").First(&userInfo).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	glog.V(3).Info("Created user successfully")
	return &userInfo, nil
}

func (u *UserDAO) GetByUsername(username string) (*do.UserResponse, error) {
	var userResponse do.UserResponse
	var user do.User
	if err := u.db.Model(&do.User{}).Where("username = ?", username).
		Preload("UserRoles.Role").
		First(&user).Error; err != nil {
		return nil, err
	}
	var profile *do.Profile
	if err := u.db.Model(&do.Profile{}).Where("user_id = ?", user.ID).Find(&profile).Error; err != nil {
		return nil, err
	}

	var Roles []string
	for _, userRole := range user.UserRoles {
		if userRole.Active {
			Roles = append(Roles, userRole.Role.Type)
		}
	}

	userResponse.ID = user.ID
	userResponse.Username = user.Username
	userResponse.FullName = profile.FirstName + " " + profile.LastName
	userResponse.Roles = Roles
	userResponse.Profile = profile

	return &userResponse, nil
}

func (u *UserDAO) UpdateUser(id int, username, password string) error {
	return nil
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db}
}
