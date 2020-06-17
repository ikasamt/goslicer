package example

import (
	"log"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	timestamp := TimeStamps{CreatedAt: time.Now(), UpdatedAt: time.Now()}
	users := Users{
		User{1, 19, `taro`, `yamada`, Man, timestamp},
		User{2, 29, `jiro`, `yamada`, Man, timestamp},
		User{3, 15, `hanako`, `yamada`, Women, timestamp},
	}

	log.Println(users.Count())

	aUser, found := users.Where(func(x User)bool{return x.Age < 20}).First()
	log.Println(found)
	log.Println(aUser)

	teens := users.Where(func(x User) bool { return x.Age < 20 })
	senior, found := teens.SortBy(func(a, b User) bool { return a.Age > b.Age }).First()
	log.Println(found)
	log.Println(senior)
	log.Println(senior.Age)

	teenFirstNames := teens.Select(`FirstName`)
	log.Println(teenFirstNames.AsString())

	teenFirstIDs := teens.Select(`ID`)
	log.Println(teenFirstIDs.AsInt())
}

