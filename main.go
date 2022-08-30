package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "sync"
	"time"

	"github.com/gorilla/mux"
)

//var websites = [3]string{"https://www.google.com", "https://www.facebook.com", "https://www.codechef.com"}

type msg struct {
	Messsage string
}

var mp = make(map[string]string)

// var wg = new(sync.WaitGroup)

// func fillmap() {
// 	for i := 0; i < len(websites); i++ {
// 		mp[websites[i]] = "Error"
// 	}
// }

func check_one(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Web_name := "https://" + params["name"]
	_, flag := mp[Web_name]
	if !flag {
		m := msg{messsage: "Website is not present"}
		json.NewEncoder(w).Encode(m)
		return
	}
	temp_map := make(map[string]string)
	temp_map[Web_name] = mp[Web_name]
	json.NewEncoder(w).Encode(temp_map)

}

// var websites = []string{}

func post_method(w http.ResponseWriter, r *http.Request) {
	temp_map := make(map[string][]string)

	err := json.NewDecoder(r.Body).Decode(&temp_map)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range temp_map["web"] {
		mp[val] = "DOWN!"
	}

}

func get_method(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&mp)
}

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {

// 		fmt.Println(link, "Down!")
// 		c <- link
// 		return
// 	}
// 	fmt.Println(link, "Up!")
// 	c <- link
// }

func check() {
	// w.Header().Set("Content-Type", "applicaton/json")
	// fmt.Fprint(w, "Welcome to server")

	for {
		for link, _ := range mp {
			resp, err := http.Get(link)
			if err != nil {
				//TODO: add a log to mention what was the error that occured
				log.Println("Error is : ", err)
				continue
			}

			if resp.StatusCode != http.StatusOK {
				mp[link] = "Down!"
				continue
			}

			mp[link] = "Up!"
		}
		time.Sleep(1 * time.Minute)
	}

	// c := make(chan string)
	// for _, link := range websites {
	// 	go checkLink(link, c)
	// }
	// for l := range c {
	// 	go func(link string) {
	// 		time.Sleep(60 * time.Second)
	// 		checkLink(link, c)
	// 	}(l)

	// }

	// get_method(w, r)

}

func main() {

	// fillmap()
	go check()

	r := mux.NewRouter()

	fmt.Println("starting server")

	// http.HandleFunc("/websites", check)
	r.HandleFunc("/websites", post_method).Methods("POST")

	r.HandleFunc("/getWebsites", get_method).Methods("GET")

	r.HandleFunc("/getWebsites/{name}", check_one).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))

	// http.HandleFunc("/websites/{name}", check_one)

	// http.ListenAndServe("localhost:3000", nil)

	// for _, val := range mp {
	// 	fmt.Println(val)
	// }

}
