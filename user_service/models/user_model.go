package models

import pb "github.com/MiracleCanCode/trello_protos/pkg/api"

func MapToModelUser(in *pb.User) *User {
	return &User{
		Id:       in.Id,
		Name:     in.Name,
		Login:    in.Login,
		Password: in.Password,
	}
}

func MapToGRPCUser(in *User) *pb.User {
	return &pb.User{
		Id:       in.Id,
		Login:    in.Login,
		Name:     in.Name,
		Password: in.Password,
		Avatar:   in.Avatar,
	}
}

type User struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Avatar   string `json:"avatar" db:"avatar"`
}
