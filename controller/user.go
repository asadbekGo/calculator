package controller

import (
	"app/models"
	"log"

	"github.com/bxcodec/faker/v4"
)

func (c *Controller) UserGenerate(count int) []*models.User {
	var users []*models.User
	for count > 0 {
		users = append(users, &models.User{
			Name:        faker.Name(),
			PhoneNumber: faker.Phonenumber(),
		})
		count--
	}

	return users
}

func (c *Controller) UserGetList(req *models.UserGetListRequest) *models.UserGetListResponse {

	log.Printf("UserGetList req: %+v\n", req)

	var (
		response = &models.UserGetListResponse{}
		offset   = c.Cfg.DefaultOffset
		limit    = c.Cfg.DefaultLimit
	)

	if req.Offset > 0 {
		offset = req.Offset
	}

	if req.Limit > 0 {
		limit = req.Limit
	}

	response.Count = len(c.Users)
	if len(c.Users) < limit+offset {
		response.Users = c.Users[offset:]
	} else {
		response.Users = c.Users[offset : offset+limit]
	}

	return response
}
