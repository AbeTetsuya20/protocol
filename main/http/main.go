package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	if r.Method == "GET" {
		t, _ := template.ParseFiles("main/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		//ログインデータがリクエストされ、ログインのロジック判断が実行されます。
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		response := "username:" + r.Form["username"][0] + " password:" + r.Form["password"][0]
		w.Write([]byte(response))
	}
}

func main() {
	http.HandleFunc("/login", login) //アクセスのルーティングを設定します

	err := http.ListenAndServe(":50001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
