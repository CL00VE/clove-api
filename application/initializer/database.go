// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package initializer

import (
	communityEnt "clove-api/domain/community/domain/ent"
	imageEnt "clove-api/domain/image/domain/ent"
	"fmt"

	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() Database {
	database := Database{
		username: os.Getenv("DATABASE_USERNAME"),
		password: os.Getenv("DATABASE_PASSWORD"),
		host:     os.Getenv("DATABASE_HOST"),
		port:     os.Getenv("DATABASE_PORT"),
		name:     os.Getenv("DATABASE_NAME"),
	}
	return database
}

type Database struct {
	username string
	password string
	host     string
	port     string
	name     string
}

// The GetUrl function is a function that returns the database url.
func (db *Database) GetUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db.username, db.password, db.host, db.port, db.name)
}

// The NewCommunityDatabase function is the function that creates the Community database.
func (db *Database) NewCommunityDatabase() (*communityEnt.Client, error) {
	client, err := communityEnt.Open("mysql", db.GetUrl())
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

// The NewImageDatabase function is the function that creates the Image database.
func (db *Database) NewImageDatabase() (*imageEnt.Client, error) {
	client, err := imageEnt.Open("mysql", db.GetUrl())
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
