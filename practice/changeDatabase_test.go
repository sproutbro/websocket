package practice

import (
	"log"
	"testing"
)

type User struct {
	Id   int
	Name string
}

type RealDB struct{}

func (RealDB) FindUser(id int) User {
	return User{Name: "실제 상용자"}
}

type FakeDB struct{}

func (FakeDB) FindUser(id int) User {
	return User{Name: "테스트 사용자"}
}

type UserFinder interface {
	FindUser(id int) User
}

func greetUser(finder UserFinder) {
	user := finder.FindUser(1)
	log.Println("안녕", user.Name)
}

func TestChangeData(t *testing.T) {

	log.Println("\n햬햬햬")

	greetUser(RealDB{})
	greetUser(FakeDB{})

	value := ChangeData()
	if value != "dddd" {
		t.Errorf("아마 맞을까??")
	}
}
