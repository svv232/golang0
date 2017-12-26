package main

import (
"fmt"
//"sync"
//"crypto/sha1"
)

/*************************************************/
var stream chan string = make(chan string, 1)

type Session struct {
    users map[string]*user
}

type user struct{
    id, authcode int
    name string
}
/*************************************************/
func (u user) chat(message string) string{
    return u.name + " :" + message
}
/*************************************************/
func newSess() *Session{
    var sess Session
    sess.users = make(map[string]*user)
    return &sess
}

func (s *Session) addUser(name string){
    authcode := 1234
    s.users[name] = &user{len(s.users), authcode, name}
}

func (s *Session) userChat(message string, name string) {
    if _, ok := s.users["name"]; ok{
        stream <- s.users[name].chat(message)
    }
}

/*************************************************/
func main(){
    sess := newSess()
    names := [4]string{"sai","tnek","captiosus", "_passion"}
    for i := range names{
        sess.addUser(names[i])
    }

    phrases := [4]string{"zippo!", "twits", "tits", "heyooo"}
    for i:=0; i < len(phrases); i++{
        sess.addUser(names[i])
        sess.userChat(phrases[i], names[i])
    }

    msg := <-stream
    close(stream)
    fmt.Scanln(msg)
}
