package myhttpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func CreateHttpServer(serverPort string) {
	file, err := openLogFile("./mylog.log")
	if err != nil {
		log.Fatal(err)
	}
	// set log to file
	log.SetOutput(file)
	// set log flags, LstdFlags = Ldate | Ltime
	// Lshortfile => apiserver.go
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	// set port number
	portNumber := strings.Join([]string{"localhost:", serverPort}, "")
	log.Print("server port : " + portNumber)

	// url
	str := "http://unknowhost.com/person.php?first=george&last=washington&age=100"
	u, _ := url.Parse(str)
	fmt.Println("original:", u)
	q, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("BEFORE")
	for key, value := range q {
		fmt.Println(key, ":", value)
	}
	q.Del("age")
	fmt.Println("AFTER")
	for key, value := range q {
		fmt.Println(key, ":", value)
	}
	u.RawQuery = q.Encode()
	fmt.Println("modified:", u)

	// router
	http.HandleFunc("/users", func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // get, select
			json.NewEncoder(wr).Encode(users)
		case http.MethodPost: // post, update
			var user User
			json.NewDecoder(r.Body).Decode(&user)
			users[user.Email] = &user
			json.NewEncoder(wr).Encode(user)
		case http.MethodDelete:
			var user User
			json.NewDecoder(r.Body).Decode(&user)
		}
	})

	// listen server, port opened
	log.Fatal(http.ListenAndServe(portNumber, nil))
}

// log file open
func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
