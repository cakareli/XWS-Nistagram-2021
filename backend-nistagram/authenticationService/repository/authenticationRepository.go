package repository

import (
	"fmt"
)

type AuthenticationRepository struct {

}

func (repository *AuthenticationRepository) Hello (){
	fmt.Printf("Hello from Repository");
}