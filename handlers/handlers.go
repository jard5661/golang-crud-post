package handlers

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:5004")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/article/", CreateArticle).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/article/", GetArticleByStatus).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/article/{limit}/{offset}", GetAllArticles).Methods("OPTIONS", "GET") 
	myRouter.HandleFunc("/article/{id}", GetArticleById).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/article/{id}", UpdateArticle).Methods("OPTIONS", "PUT")
	myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("OPTIONS", "POST")
	
	handler := cors.AllowAll().Handler(myRouter)
	/* log.Fatal(http.ListenAndServe(handler) */
	log.Fatal(http.ListenAndServe(":5004", handler))
}