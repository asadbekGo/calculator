package models

type User struct {
	Name        string
	PhoneNumber string
}

type UserGetListRequest struct {
	Offset int
	Limit  int
}

type UserGetListResponse struct {
	Count int
	Users []*User
}
