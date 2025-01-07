package models

import (
	//"database/sql"
	"context"
	//"encoding/json"
	//"fmt"
	svc "gisapi/database"
	//"gisapi/dto"
	"github.com/jackc/pgx/v5"
	"strings"
)

// DAO
type User struct {
	ID       int
	UserName string
	Hash     string
	UUID     string
}

func GetUserFromUsername(username string) (User, error) {

	//const sql= `SELECT row_to_json(f) as user FROM (SELECT uuid::text as uuid,concat(trim(lastname),' ',(firstname)) as username,trim(mail) as mail FROM  users WHERE trim(mail)=trim(@mail)) f;`

	const sql = `SELECT id,concat(trim(lastname),' ',(firstname)) as username,trim(hash),uuid::text as uuid FROM  users  WHERE trim(mail)=@mail;`

	args := pgx.NamedArgs{"mail": strings.TrimSpace(username)}

	row := svc.DB.QueryRow(context.Background(), sql, args)

	//fmt.Println("Row:", row)

	var user User

	err := row.Scan(&user.ID, &user.UserName, &user.Hash, &user.UUID)

	if err != nil {
		//fmt.Println("err:", err)
		return User{}, err
	}
	//fmt.Println("User:", user)
	return user, nil

}

func GetUserIdFromUUID(uuid string) (int, error) {

	const sql = `SELECT id FROM  users  WHERE Trim(uuid)=Trim(@uuid);`

	args := pgx.NamedArgs{"uuid": uuid}

	row := svc.DB.QueryRow(context.Background(), sql, args)

	//fmt.Println("Row:", row)

	var id int

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil

}
