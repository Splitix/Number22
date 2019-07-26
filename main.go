package main

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"
)

type Todo struct {
    Name      string
}

type Todos []Todo

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  todos := Todos{
      Todo{Name: "Write presentation"},
      Todo{Name: "Host meetup"},
  }

  json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}


// package main
//
// import (
//   "fmt"
//   "net/http"
// )
//
// // type player struct {
// //   currentGold int
// //   maxGold  int
// //   wins int
// //   loses int
// // }
//
// func main() {
//   http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Welcome to my website!")
// 	})
//
// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
//
// 	http.ListenAndServe(":80", nil)
// 	// fmt.Println("Welcome to 22")
//   // fmt.Println("Select Difficulty:")
//   // fmt.Println("-- 1) Beginner [300 Gold]")
//   // fmt.Println("-- 2) Moderate [500 Gold]")
//   // fmt.Println("-- 3) Hard     [1000 Gold]")
//   // fmt.Println("-- 4) God Mode [2000 Gold]")
//   //
//   // var userOptionChoice int
//   // fmt.Print("-> ")
//   // fmt.Scan(&userOptionChoice)
//   //
//   // user := player{currentGold: 100, wins: 0, loses: 0}
//   //
//   // switch userOptionChoice {
//   // case 1:
//   //   fmt.Println("You selected 1")
//   //   user.maxGold = 300
//   // case 2:
//   //   fmt.Println("You selected 2")
//   //   user.maxGold = 500
//   // case 3:
//   //   fmt.Println("You selected 3")
//   //   user.maxGold = 1000
//   // case 4:
//   //   fmt.Println("You selected 4")
//   //   user.maxGold = 2000
//   // default:
//   //   fmt.Println("Not a valid option please try again")
//   // }
//   //
//   // fmt.Println("Player 1 has", user.getUserGold())
// }
//
// // func (p player) getUserGold() int{
// //   return p.currentGold
// // }
