package player

import "fmt"

type Avatar struct {
	Url string
}

type Player struct {
	Name string
	Age int
	Avatar
	password string
}

type User struct {
	Name string
}

func (u User) SayHello() {
	u.Name = "Va te faire ... par les grecs"
	fmt.Printf("Hello %s\n", u.Name)
}


func New(name string) Player  {
	return Player{
		Name:     name,
		password: "default",
	}
}
