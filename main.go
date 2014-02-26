package main

import (
	"fmt"
	"github.com/lunny/godbc"
	 "github.com/lunny/xorm"
)

type User struct { Id int64 Name string `xorm:"varchar(25) not null unique 'usr_name'"` } 

func main() {

	engine, err = xorm.NewEngine("MsSql", "SERVER=127.0.0.1;UID=sa;PWD=123;DATABASE=DemoDB")
	defer engine.Close()

	user := new(User)
user.Name = "myname"
affcted, err := engine.Insert(user)
}
