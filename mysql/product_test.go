package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindAllProductRepo(t *testing.T) {
	testdb, err := ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testdb.Close()
	r, _ := NewProductRepository(testdb)

	resp, _ := r.FindAllProductRepo()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, resp)
}

func Test_FindProductByID(t *testing.T) {
	testdb, err := ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testdb.Close()
	r, _ := NewProductRepository(testdb)

	resp, _ := r.FindProductByIDRepo(1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, resp)
}

func Test_FavoriteProduct(t *testing.T) {
	testdb, err := ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testdb.Close()
	r, _ := NewProductRepository(testdb)

	resp, _ := r.FavoriteProductRepo()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, resp)
}

func Test_FindCategoryByCode(t *testing.T) {
	testdb, err := ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testdb.Close()
	r, _ := NewProductRepository(testdb)

	resp, _ := r.FindCategoryByCodeRepo("01")
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, resp)
}

func Test_FindCategory(t *testing.T) {
	testdb, err := ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testdb.Close()
	r, _ := NewProductRepository(testdb)

	resp, _ := r.FindCategoryRepo()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, resp)
}

func Test_FindPackage(t *testing.T) {
	testdb, err := ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testdb.Close()
	r, _ := NewProductRepository(testdb)

	resp, _ := r.FindPackageRepo()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, resp)
}
