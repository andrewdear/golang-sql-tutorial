package main

import (
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var PostgresPool *sqlx.DB

func ConnectToPostgres() error {
	// Make not that we are only using a plain string here as this is local development.
	// For use in a production system you should ensure these credentials are stored somewhere like aws secrets or parameter store or the authentication is handled by ssh keys.
	// Say can creat tutorial if anyone wants that too.
	dbPool, err := sqlx.Open("postgres", "postgres://postgres:secret@localhost:5432/golang_sql?sslmode=disable")

	if err != nil {
		return err
	}

	PostgresPool = dbPool

	// check that connection has worked
	if err := PostgresPool.Ping(); err != nil {
		return err
	}

	return nil
}

func GetPostgresPool() *sqlx.DB {
	if PostgresPool == nil {
		err := ConnectToPostgres()
		if err != nil {
			log.Fatal(err, "Initial database connection has failed, check credentials")
		}
	}

	if err := PostgresPool.Ping(); err != nil {

		err = ConnectToPostgres()
		// Keep trying to connect to postgres
		if err != nil {
			log.Error(err, "Failed to get postgres connection, retrying")
			return GetPostgresPool()
		}
	}

	return PostgresPool
}

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

type Todo struct {
	ID      int    `db:"id"`
	Content string `db:"content"`
	UserId  string `db:"user_id"`
}

func InsertUsersWithStructs() error {
	users := []User{
		{Username: "captain_sparkles", Email: "captain@sparkles.com"},
		{Username: "sergeant_sparkles", Email: "sergeant@sparkles.com"},
	}
	// REMEMBER: say a names exec is alot easier than a names query, we will go through those in a little while.
	_, err := GetPostgresPool().NamedExec("INSERT INTO users (username, email) VALUES (:username, :email)", users)

	return err
}

// Maps are often used for things like join tables ect if you have structured data then use structs.
func InsertTodosWithMaps() error {
	todos := []map[string]interface{}{
		{"content": "steer the ship", "user_id": 1},
		{"content": "command the crew", "user_id": 1},
		{"content": "Man the cannons", "user_id": 2},
	}

	_, err := GetPostgresPool().NamedExec("INSERT INTO todos (content, user_id) VALUES (:content, :user_id)", todos)

	return err
}

func GetMultipleTodos() (error, *[]Todo) {
	// Show ANY talk about how it knows how to handle slices ect so really easy to use
	var todos []Todo

	err := GetPostgresPool().Select(&todos, "SELECT * FROM todos WHERE user_id = ANY($1)", []int{1, 2})

	return err, &todos
}

func (user *User) GetUser() error {
	return GetPostgresPool().Get(user, `SELECT * FROM users WHERE id=$1`, user.ID)
}

func NamedExternalScriptQuery(userId int) (error, *[]Todo) {
	sql, err := os.ReadFile("sql_templates/get_user.sql")
	if err != nil {
		return err, nil
	}

	rows, err := GetPostgresPool().NamedQuery(string(sql), map[string]interface{}{"user_id": userId})

	if err != nil {
		return err, nil
	}

	var todos []Todo

	if rows != nil {
		for rows.Next() {
			var todo Todo
			err = rows.StructScan(&todo)
			if err != nil {
				return err, nil
			}

			todos = append(todos, todo)
		}

	}

	return err, &todos
}

func GetSingleField() (error, []int) {
	var todoIds []int

	err := GetPostgresPool().Select(&todoIds, "SELECT id FROM todos")

	return err, todoIds
}

func DeleteRecord(todoId int) error {

	query := `DELETE FROM todos WHERE id = $1;`
	// Sometimes normal base commands are best. This library gives you that option
	_, err := GetPostgresPool().Exec(query, todoId)

	if err != nil {
		return errors.New("failed to delete Todo")
	}

	return nil
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	err := GetPostgresPool().Ping()
	defer GetPostgresPool().Close()

	if err != nil {
		log.Error("ugh oh")
	}

	log.Info("Connected to postgres!")

	//err = InsertUsersWithStructs()
	//
	//if err != nil {
	//	log.Error(err, "Failed to insert users with structs")
	//}
	//
	//log.Info("Inserted users with structs")

	//err = InsertTodosWithMaps()
	//
	//if err != nil {
	//	log.Error("failed to insert todos")
	//}
	//
	//log.Info("Inserted todos with maps")

	//user := &User{ID: 2}
	//err = user.GetUser()
	//if err != nil {
	//	log.Error("failed to get user")
	//}
	//
	//log.Infof("%+v", *user)

	//err, todos := GetMultipleTodos()
	//
	//if err != nil {
	//	log.Error("failed to get todos")
	//}
	//
	//log.Infof("%+v", *todos)

	//err, todosIds := GetSingleField()
	//
	//if err != nil {
	//	log.Error(err, "Failed to get todos")
	//}
	//
	//log.Info(todosIds)

	//err, todos := NamedExternalScriptQuery(1)
	//
	//if err != nil {
	//	log.Error(err, "Failed to get todos")
	//}
	//
	//if todos != nil {
	//	log.Infof("%+v", *todos)
	//}

	err = DeleteRecord(3)

	if err != nil {
		log.Error(err, "Failed to get todos")
	}

}
