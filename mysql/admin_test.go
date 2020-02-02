package mysql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addAdmin(t *testing.T) {
	testdb, _ := ConnectDB()
	defer testdb.Close()

	x, _ := NewAdminRepository(testdb)
	resp, err := x.AddUserAdminRepo("bn620900240001", "nziomini@gmail.com", "$2a$08$biWY.McVnnyqVlFtSxJO2Ox.WrbCir5.WQTjhdypbcOOCBmqAwpPC")
	fmt.Println(resp)
	if err != nil {
		t.Fatalf("error get BranchId : %v", err.Error())
	}

	assert.Equal(t, resp, "บริษัท นพดลพานิช จำกัด (สำนักงานใหญ่)")
}
