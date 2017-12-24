package main

import (
"fmt"
)

type user struct{
    id, authcode int
    name string
}

type Session struct {
    users map[string]*user
}

func (u user) chat(message string){
    fmt.Println(u.name + " :" + message)
}

func newSess() *Session{
    var sess Session
    sess.users = make(map[string]*user)
    return &sess
}

func (s Session) addUser(name string){
    authcode := 1234
    s.users[name] = &user{len(s.users), authcode, name}
}

func main(){
    sess := newSess()
    sess.addUser("sai")
    sess.addUser("tnek")
    sess.users["sai"].chat("hey doo")
    sess.users["tnek"].chat("WazZoop!")
}
