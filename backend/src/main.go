package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"


	_ "github.com/go-sql-driver/mysql"
)

type EventData struct {
	ID   int    `json:"id"`
	EventContent string `json:"event_content"`
}
func main() {
	//define router

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

	if err = recordStats(db, 1 /*some event id*/, "get some milk"); err != nil {
		panic(err)
	}

	fmt.Println("pass the test successfully for db connection")

	//router
	r := mux.NewRouter()

	fmt.Println("pass the test successfully for api connection")

	//api endpoints
	r.Handle("/todo", http.HandlerFunc(ToDoCreate)).Methods("POST")
	r.Handle("/todolist", http.HandlerFunc(TodoList)).Methods("GET")

	//define options
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	//start server
	log.Fatal(http.ListenAndServe(":8081", corsWrapper.Handler(r)))
}


func recordStats(db *sql.DB, eventID int64, eventName string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE event SET event_content = 'updated event content' where id = $1",eventID); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO event (event_content) VALUES ($1)", eventName); err != nil {
		return
	}
	return
}

func todoList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		respData := make([]*EventData, 0)
		for rows.Next() {
			resp := new(EventData)
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

		for i:= 0; i< len(respData); i++{
			data:=fmt.Sprintf("todolist event content: %s",respData[i].EventContent)
			println(data)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(respData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading bookmarks"))
		}
	})
}

func todoCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		sq:= fmt.Sprintf(`insert into event (event_content) values ('$1')`)
		_,err = db.Exec(sq,reqBody.EventContent)
		if err != nil {
			response(nil, 500, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(reqBody); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading bookmark"))
		}
	})
}
