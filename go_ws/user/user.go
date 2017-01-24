package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID    int    "json:user_id"
	Name  string "json:user_nickname"
	Email string "json:user_email"
	First string "json:user_first"
	Last  string "json:user_last"
}

type CreateResponse struct {
	Error     string "json:error"
	ErrorCode int    "json:code"
}

type ErrMsg struct {
	ErrCode    int
	StatusCode int
	Msg        string
}

func ErrorMessages(err int64) ErrMsg {
	em := ErrMsg{}
	errorMessage := ""
	statusCode := 200
	errorCode := 0
	switch err {
	case 1062:
		errorMessage = "Duplicate entry"
		errorCode = 10
		statusCode = 409
	}
	em.ErrCode = errorCode
	em.StatusCode = statusCode
	em.Msg = errorMessage
	return em
}
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("interface.html")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(404)
		fmt.Fprintf(w, "Page cannot loaded!")
	}
	t.Execute(w, nil)
}

var db *sql.DB

func initMysql() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/social_network?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
}
func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := User{}
	user.Name = r.FormValue("user")
	user.Email = r.FormValue("email")
	user.First = r.FormValue("first")
	user.Last = r.FormValue("last")
	f, _, err := r.FormFile("image1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileData, _ := ioutil.ReadAll(f)
	fileString := base64.StdEncoding.EncodeToString(fileData)
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		log.Println("Json parsing error!")
		w.WriteHeader(502)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Println(string(output))
	response := CreateResponse{}
	//insert into the database
	sql := "INSERT INTO users SET user_image=?,user_nickname=?,user_email=?,user_first=?,user_last=?"
	result, err := db.Exec(sql, fileString, user.Name, user.Email, user.First, user.Last)
	if err != nil {
		errorMsg, errorCode := dbErrorParse(err.Error())
		e := ErrorMessages(errorCode)
		fmt.Println(errorMsg)
		log.Println(err.Error())
		response.Error = e.Msg
		response.ErrorCode = e.ErrCode
		http.Error(w, "Conflict", e.StatusCode)
	}
	fmt.Println(result)
	output, _ = json.MarshalIndent(&response, "", "\t")
	w.WriteHeader(200)
	w.Write(output)
}

func dbErrorParse(err string) (string, int64) {
	Parts := strings.Split(err, ":")
	errorMessage := Parts[1]
	Code := strings.Split(Parts[0], "Error ")
	errorCode, _ := strconv.ParseInt(Code[1], 10, 32)
	return errorMessage, errorCode
}

func UserRetrieve(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	id := vals["id"]
	nid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		log.Fatal(err)
		fmt.Fprintf(w, "Invalid id!")
	}
	user := User{}
	err = db.QueryRow("SELECT user_id,user_nickname,user_email,user_first,user_last FROM users where user_id=?", nid).Scan(&user.ID, &user.Name, &user.Email, &user.First, &user.Last)
	switch {
	case err == sql.ErrNoRows:
		w.Header().Set("Pragma", "No-Cache")
		w.WriteHeader(200)
		fmt.Fprintf(w, "No such user")
	case err != nil:
		log.Fatal(err)
		fmt.Fprintf(w, "Error")
	default:
		b, err := json.MarshalIndent(&user, "", "\t")
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(w, "Something error!")
			return
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func UsersRetrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Pragma", "No-Cache")
	w.Header().Set("Content-Type", "application/json")
	rows, _ := db.Query("SELECT user_id,user_nickname,user_email,user_first,user_last FROM users")
	data := make([]User, 0)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.First, &user.Last)
		data = append(data, user)
	}
	b, err := json.Marshal(&data)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprintf(w, "Something error!")
		return
	}
	w.Write(b)
}

func main() {
	initMysql()
	router := mux.NewRouter()
	router.HandleFunc("/api/users", UserCreate).Methods("POST")
	router.HandleFunc("/api/users", UsersRetrieve).Methods("GET")
	router.HandleFunc("/api/user/read/{id:[0-9]+}", UserRetrieve).Methods("GET")
	router.HandleFunc("/", IndexHandle)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
