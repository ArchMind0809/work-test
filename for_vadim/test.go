package main

import (
	"fmt"
	"net/http"
	//"strconv"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
type product struct{
    id int
    name string
}

func main() {


	db, err := sql.Open("sqlite3", "filedb.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Get_f := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		input_var := query.Get("name")
		result := db.QueryRow("select * from emptyFiles where Name = $1", input_var)
		prod = []product{}
		for result.Next() {
        	result.Scan(&id, &name)
        	fmt.Println(strconv.Itoa(id) + ": " + name)
    	}
		if err != nil{
    		panic(err)
		}
		Get_Func(result)
	}

	Post_f := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		input_var := query.Get("name")
		result, err := db.Exec("insert into emptyFiles (Name) values ($1)", input_var)
		if err != nil{
    		panic(err)
		}
		//Post_Func(result)
		fmt.Println(result)
	}

	hi := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi")
	}

	//http.HandleFunc("/get", Get_f)
	http.HandleFunc("/post", Post_f)
	http.HandleFunc("/", hi)
	
	
	fmt.Println("Server is listening...")
    http.ListenAndServe(":80", nil)
}



// func Get_Func (result string) {
// 	prod := product{}
// 	err = row.Scan(&prod.id, &prod.name)
// 	if err != nil{
//     	panic(err)
// 	}
// 	fmt.Println(prod.id, prod.name)
// 	return 0
//}
// func Post_Func (input_var string) {
// 	return 0
// }
