package service

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/repository"
)

type FollowService struct {
	FollowRepository *repository.FollowRepository
}
