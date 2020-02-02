package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	gEnv     = "development" //default
	gSSLMode = "disable"
	bPort    = "3306"
	myHost   = "perfect_db.extensionsoft.biz"
	myUser   = "perfect"
	myDb     = "perfect"
	myPass   = "P@ssw0rd"
)

func ConnectDB() (db *sql.DB, err error) {
	dsn := myUser + ":" + myPass + "@tcp(" + myHost + ":" + bPort + ")/" + myDb + "?parseTime=true&charset=utf8&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql error =", err)
		return nil, err
	} else {
		log.Println("connect sucess", dsn)
	}
	db.Exec("use " + myDb)

	return db, nil
}

func Test_getInviteCode(t *testing.T) {
	testdb, _ := ConnectDB()
	defer testdb.Close()

	x, _ := NewAuthRepository(testdb)
	resp := x.GetInviteCode()
	fmt.Println(resp)

	assert.Equal(t, resp, "บริษัท นพดลพานิช จำกัด (สำนักงานใหญ่)")
}
func Test_SearchBranchById(t *testing.T) {
	testdb, _ := ConnectDB()
	defer testdb.Close()

	x, _ := NewAuthRepository(testdb)
	resp, err := x.FindUserID()
	fmt.Println(resp)
	if err != nil {
		t.Fatalf("error get BranchId : %v", err.Error())
	}

	assert.Equal(t, resp, "บริษัท นพดลพานิช จำกัด (สำนักงานใหญ่)")
}
func Test_FindUserTel(t *testing.T) {
	testdb, _ := ConnectDB()
	defer testdb.Close()

	x, _ := NewAuthRepository(testdb)
	resp, err := x.FindUserTel("+66816398388")
	fmt.Println(resp)
	if err != nil {
		t.Fatalf("error get BranchId : %v", err.Error())
	}

	assert.Equal(t, resp, "บริษัท นพดลพานิช จำกัด (สำนักงานใหญ่)")
}

func Test_FindUserEmail(t *testing.T) {
	testdb, _ := ConnectDB()
	defer testdb.Close()

	x, _ := NewAuthRepository(testdb)
	resp, err := x.FindUserEmail("nziomini@gmail.com")
	fmt.Println(resp)
	if err != nil {
		t.Fatalf("error get BranchId : %v", err.Error())
	}

	assert.Equal(t, resp, "บริษัท นพดลพานิช จำกัด (สำนักงานใหญ่)")
}
