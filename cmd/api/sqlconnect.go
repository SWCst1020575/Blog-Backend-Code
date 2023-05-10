package api

import (
	"database/sql"
	"fmt"

	"context"
	"net/http"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

//var TodoList []Todo

/* const (
	// db info
	HOST     = "127.0.0.1"
	PORT     = "5432"
	DATABASE = "postgres"
	USER     = "mypostgres"
	PASSWORD = "pass"
)
*/
type Todo struct {
	Id   int64
	Item string
}
type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	Tag         string `json:"tag"`
	Link        string `json:"link"`
	ImgLink     string `json:"imgLink"`
}

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

/*
func addProject(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 4096))

	checkRequestError(err, w, r)
	var newProject Project
	_ = json.Unmarshal(body, &newProject)
	defer r.Body.Close()
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE),
	)
	checkServerError(err, w, r)
	defer db.Close()
	res, err := db.Exec("INSERT INTO project VALUES (DEFAULT, $1, $2, $3, $4, $5, $6);",
		newProject.Title, newProject.Description, newProject.Lang, newProject.Tag,
		newProject.Link, newProject.ImgLink)
	checkServerError(err, w, r)
	fmt.Println(res)

	response := ApiResponse{"200", newProject}

	ResponseWithJson(w, http.StatusOK, response)
} */
func getProjectAll(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlserver",
		fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			server, user, password, port, database))
	checkServerError(err, w, r)

	ctx := context.Background()
	err = db.PingContext(ctx)
	checkServerError(err, w, r)
	fmt.Println("Connected!")
	tsql := fmt.Sprintf("SELECT * FROM project;")

	rows, err := db.QueryContext(ctx, tsql)
	defer rows.Close()
	checkServerError(err, w, r)
	var newProjectList []Project
	for rows.Next() {
		var id int
		var newProject Project
		err = rows.Scan(&id, &newProject.Title, &newProject.Description, &newProject.Lang, &newProject.Tag, &newProject.Link, &newProject.ImgLink)
		newProjectList = append(newProjectList, newProject)
		checkServerError(err, w, r)
	}
	if len(newProjectList) == 0 {
		checkRequestNotfound(w, r)
		return
	}
	response := ApiResponse{"200", newProjectList}
	ResponseWithJson(w, http.StatusOK, response)
}

/* func getProjectByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"]
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE),
	)
	checkServerError(err, w, r)
	defer db.Close()
	row, err := db.Query("SELECT * FROM project where id=$1", queryId)
	defer row.Close()
	checkServerError(err, w, r)
	var id int
	var newProject Project
	for row.Next() {
		err = row.Scan(&id, &newProject.Title, &newProject.Description, &newProject.Lang, &newProject.Tag, &newProject.Link, &newProject.ImgLink)
		checkServerError(err, w, r)
	}
	if id == 0 {
		checkRequestNotfound(w, r)
		return
	}
	response := ApiResponse{"200", newProject}
	ResponseWithJson(w, http.StatusOK, response)
}
*/
