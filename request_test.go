package request

import (
	"log"
	"net/url"
	"strconv"
	"testing"
)

func TestGetAll(t *testing.T) {
	getPosts, err := get("/posts", nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(getPosts))
}

func TestPost(t *testing.T) {
	postBody := map[string]string{
		"title":  "foo",
		"body":   "bar",
		"userId": "1",
	}

	createPost, err := post("/posts", postBody)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(createPost))
}
func TestGetOneWithId(t *testing.T) {
	for i := 0; i < 10; i++ {
		getOne, err := get("/posts/"+strconv.Itoa(i), nil)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(getOne))
	}
}
func TestGetOneWithParam(t *testing.T) {
	for i := 0; i < 10; i++ {
		getOne, err := get("/posts", url.Values{
			"id": {strconv.Itoa(i)},
		})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(getOne))
	}
}
