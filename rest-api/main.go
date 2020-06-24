package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is ..
type Article struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles []Article

func initializeArticles() {
	articles = []Article{
		{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range articles {
		if article.ID == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleArticleRequests(router *mux.Router) {
	router.HandleFunc("/articles", returnAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", returnSingleArticle)
	router.HandleFunc("/articles", createNewArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	handleArticleRequests(router)

	log.Println(http.ListenAndServe(":10000", router))
}

func main() {
	initializeArticles()
	handleRequests()
}
