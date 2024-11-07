package rpc

import (
	"context"
	"fmt"

	"github.com/MarskTM/financial_report_server/infrastructure/database/do"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/MarskTM/financial_report_server/utils"
	"github.com/golang/glog"
)

func (s *BizService) Authenticate(ctx context.Context, req *pb.Credentials) (*pb.AuthResponse, error) {
	if req.Username == "" {
		err := fmt.Errorf("Error not found username for Authenticate!")
		glog.V(1).Info(err)
		return nil, err
	}

	user, err := s.BizModel.DB.UserDAO.GetByUsername(req.Username)
	if err != nil {
		glog.V(1).Infof("Error getting user by username: %v", err)
		return nil, err
	}

	// Gen sessions id
	// ...

	return &pb.AuthResponse{
		Session:   1,
		UserId:    user.ID,
		Usernames: user.FullName,
		Roles:     user.Roles,
	}, nil
}

func (s *BizService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Username == "" {
		err := fmt.Errorf("Error not found username for RegisterClient!")
		glog.V(1).Info(err)
		return nil, err
	}

	//
	newUser := do.User{
		Username: req.Username,
		Password: utils.HashAndSalt(req.Password),
	}

	profile := do.Profile{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.PhoneNumber,
		Birthdate: req.BirthDate,
	}

	user, err := s.BizModel.DB.UserDAO.CreateUser(newUser, profile)
	if err != nil {
		glog.V(1).Infof("Error creating new user: %v", err)
		return nil, err
	}

	var Roles []string
	for _, userRole := range user.UserRoles {
		if userRole.Active {
			Roles = append(Roles, userRole.Role.Type)
		}
	}

	// Gen sessions id
	// ...

	return &pb.RegisterResponse{
		Success: true,
		Message: "Client registered successfully!",
		Auth: &pb.AuthResponse{
			Session:   1,
			UserId:    user.ID,
			Usernames: profile.FirstName + " " + profile.LastName,
			Roles:     Roles,
		},
	}, nil
}
