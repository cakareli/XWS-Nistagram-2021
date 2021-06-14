package repository

import (
	"fmt"
)

type FollowRepository struct {

}

func (repository *FollowRepository) Hello (){
	fmt.Printf("Hello from Repository")
}
