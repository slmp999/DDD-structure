package mysql

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchID(t *testing.T) {
	testdb, _ := ConnectDB()
	defer testdb.Close()

	x, _ := NewMemberRepository(testdb)
	resp, err := x.GeTProfileMemberRepo("62082400014")
	fmt.Println(resp)
	if err != nil {
		t.Fatalf("error get BranchId : %v", err.Error())
	}

	assert.Equal(t, resp, "บริษัท นพดลพานิช จำกัด (สำนักงานใหญ่)")
}

func Test_c(t *testing.T) {
	test := ConvertPhonetoCoutryCode("+6666979679089")
	randomKey := strings.Replace(test, " ", "", -1)
	fmt.Println(randomKey)
}

func Test_t(t *testing.T) {
	s := hashAndSalt("1234")
	fmt.Println(s)
}
