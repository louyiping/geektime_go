package week2

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	userId string
}

func main() {
	user, foundUser, err := GetUserDao("10086")
	if err != nil {
		fmt.Printf("FATAL error:%+w\n", err)
	}
	if !foundUser {
		fmt.Printf("the user with id %s not exist", foundUser)
	}
	fmt.Printf("the user info is %v", user)
}

func GetUserDao(userId string) (User, bool, error) {
	// execute sql to get user
	user, err := ExecuteSQL()
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false, nil
		}
		return User{}, false, errors.Wrap(err, "get user failed")
	}
	return user, true, nil
}

func ExecuteSQL() (User, error) {
	return User{}, sql.ErrNoRows
}
