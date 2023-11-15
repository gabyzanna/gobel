package gobel

import (
	"fmt"
	"testing"
	"github.com/whatsauth/watoken"
)
 
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("belcoba", privateKey)
	fmt.Println(hasil, err)
}

func TestInsertUser(t *testing.T) {
	mconn := GetConnectionMongo("MONGOSTRING", "gibel")
	var userdata User
	userdata.Username = "bella"
	userdata.Password = "cobain"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}