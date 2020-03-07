package healthcheck

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func GetUserNameFromDB(connection *sqlx.DB) (User, error) {
	user := User{}
	err := connection.Get(&user, "SELECT id,name FROM user WHERE ID=1")
	if err != nil {
		fmt.Printf("Get user name from tearup get error : %s", err.Error())
		return User{}, err
	}
	return user, nil
}
