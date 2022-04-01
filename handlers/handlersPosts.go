package handlers

import (
	"time"
	"encoding/json"
	"golang-test/connection"
	"golang-test/structs"
	"net/http"
	"github.com/thedevsaddam/govalidator"
	"github.com/gorilla/mux"

)

func CreateArticle(w http.ResponseWriter, r *http.Request){
	var post structs.Post
	rules := govalidator.MapData{
		"title": 		[]string{"required", "min:20", "max:200"},
		"content":    	[]string{"required", "min:200"},
		"category":     []string{"required", "min:3","max:100"},
		"status":      	[]string{"required"},
	}

	Messages := govalidator.MapData{
		"status":      	[]string{"required:Publish,Draft or Thrash"},
	}
	opts := govalidator.Options{
		Request: r,
		Data:    &post,
		Rules:   rules,
		Messages:    Messages,

	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}
	if len(e) != 0 {
		res := structs.Result{Code:400, Data:err, Message:"Bad Request"}
		result, err := json.Marshal(res)	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}else{
		connection.DB.Create(&post)
		res := structs.Result{Code: 200,  Message: "Success Create Article"}
		result, err := json.Marshal(res)	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
} 

func GetAllArticles(w http.ResponseWriter, r *http.Request){

	Post := []structs.Post{}
	params := mux.Vars(r)
	limit := params["limit"]
	offset := params["offset"]
	if limit == "" || offset == "" {
		limit = "10"
		offset = "0"
	} 


	connection.DB.
		Limit(limit).
		Offset(offset).
		Find(&Post)


	res := structs.Result{Code: 200, Data: Post, Message: "Get All Article, page = "+ offset +" take = "+ limit +" Success"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var post structs.Post
    connection.DB.First(&post, id)


	res := structs.Result{Code: 200, Data: post, Message: "Get Article By Id"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var post structs.Post
	
	rules := govalidator.MapData{
		"title": 		[]string{"required", "min:20", "max:200"},
		"content":    	[]string{"required", "min:200"},
		"category":     []string{"required", "min:3","max:100"},
		"status":      	[]string{"required"},
	}

	Messages := govalidator.MapData{
		"status":      	[]string{"required:Publish,Draft or Thrash"},
	}
	opts := govalidator.Options{
		Request: r,
		Data:    &post,
		Rules:   rules,
		Messages:    Messages,

	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}
	if len(e) != 0 {
		res := structs.Result{Code:400, Data:err, Message:"Bad Request"}
		result, err := json.Marshal(res)	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}else{
		update := structs.Post{Title:post.Title, Content:post.Content, Category:post.Category, Updated_date:time.Now().UTC(), Status:post.Status}
		connection.DB.First(&post, id)
		connection.DB.Model(&post).Updates(update)
		res := structs.Result{Code: 200, Message: "Success Update Article"}
		result, err := json.Marshal(res)	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post structs.Post
	delete := structs.Post{Updated_date:time.Now(), Status:"Thrash"}
	connection.DB.First(&post, id)
	connection.DB.Model(&post).Updates(delete)

	res := structs.Result{Code: 200, Message: "Success delete article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetArticleByStatus(w http.ResponseWriter, r *http.Request){
	post := []structs.Post{}

	status := r.FormValue("status")
	
	if status != "Publish" && status != "Draft" && status != "Thrash" {
		res := structs.Result{Code:400, Message:"Publish,Draft or Thrash"}
		result, err := json.Marshal(res)	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	}else {
		connection.DB.Where("status = ?",status).Find(&post)


		res := structs.Result{Code: 200, Data: post, Message: "Get All Articles Status = "+ status +""}
	
		result, err := json.Marshal(res)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	}

}