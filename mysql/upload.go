package mysql

import (
	"database/sql"
	// "fmt"

	"github.com/jmoiron/sqlx"
	//"github.com/labstack/gommon/log"

	//log "gitlab.com/satit13/perfect_api/logger"
	"gitlab.com/satit13/perfect_api/member"
	"upper.io/db.v3/lib/sqlbuilder"

	//"upper.io/db.v3/postgresql"
	"upper.io/db.v3/mysql"
)

// NewAuthRepository creates new auth repository
func NewUploadRepository(db *sql.DB) (member.Repository, error) {
	// fmt.Println(1)
	pdb, err := mysql.New(db)
	if err != nil {
		return nil, err
	}
	dbx := sqlx.NewDb(db, "mysql")
	// fmt.Println(1)
	r := memberRepository{pdb, dbx}
	// fmt.Println(1)
	return &r, nil
}

type uploadRepository struct {
	db  sqlbuilder.Database
	dbx *sqlx.DB
}
