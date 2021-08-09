package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type todoReqBody struct {
	EventContent string `json:"EventContent"`
}

type todoRespBody struct {
	ID      int64  `json:"ID"`
	EventContent string `json:"EventContent"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "tayitkan"
	dbname   = "todoapp"
)

func TodoList(w http.ResponseWriter, r *http.Request) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	sq:= fmt.Sprintf(`select id,event_content from event `)
	rows, err := db.Query(sq)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	respData := make([]*todoRespBody, 0)
	for rows.Next() {
		resp := new(todoRespBody)
		if err = rows.Scan(&resp.ID,&resp.EventContent); err != nil {
			panic(err)
		}
		respData = append(respData, resp)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	//end of api response
	err = response(respData, 200, w)
	if err != nil {
		response(nil, 500, w)
		return
	}
}

func ToDoCreate(w http.ResponseWriter, r *http.Request) {

	//get request body into loginreqbody struct
	var reqBody todoReqBody
	err := BodyToJson(r, &reqBody)
	if err != nil {
		response(nil, 500, w)
		return
	}

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	sq:= fmt.Sprintf(`insert into event (event_content) values ($1)`)
	_,err = db.Exec(sq,reqBody.EventContent)
	if err != nil {
		response(nil, 500, w)
		return
	}

	//end of api response
	err = response(reqBody, 200, w)
	if err != nil {
		response(nil, 500, w)
		return
	}
}


