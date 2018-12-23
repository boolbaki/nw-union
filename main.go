package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	//"encoding/json"

	"github.com/julienschmidt/httprouter"
	"github.com/olivere/elastic"
)

type Story struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("=== Application Starting!!")
	router := httprouter.New()
	router.GET("/hc", healthCheckHandler)
	router.GET("/stories/:id", getStory)
	router.PUT("/stories/:id", putStory)
	router.DELETE("/stories/:id", deleteStory)

	// test
	router.GET("/test", postStory)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("--- healthHandler")
	fmt.Fprint(w, "OK")
}

func getStory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("--- getStory")

	client, err := elastic.NewClient(
		elastic.SetURL("http://es:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	id := ps.ByName("id")
	resp, err := client.Get().
		Index("story").
		Type("external").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(*resp.Source)
}

func postStory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("--- postStory")
	client, err := elastic.NewClient(
		elastic.SetURL("http://es:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	//id := ps.ByName("id")
	resp, err := client.Update().
		Index("story").
		Type("external").
		Id("3").
		Doc(map[string]interface{}{"title": "update", "body": "fo"}).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Id)

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(*resp.Source)
}

func putStory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("--- putStory")
	client, err := elastic.NewClient(
		elastic.SetURL("http://es:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	id := ps.ByName("id")
	resp, err := client.Update().
		Index("story").
		Type("external").
		Id(id).
		Doc(map[string]interface{}{"title": "update", "body": "fo"}).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Id)

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(*resp.Source)
}

func deleteStory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("--- deleteStory")
	client, err := elastic.NewClient(
		elastic.SetURL("http://es:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	id := ps.ByName("id")
	resp, err := client.Delete().
		Index("story").
		Type("external").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Id)
	//w.Header().Set("Content-Type", "application/json")
	//w.Write("deleted")
	fmt.Fprint(w, "deleted")
}
