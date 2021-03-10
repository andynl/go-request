package request

import (
	"log"
	"net/url"
	"strconv"
	"testing"
)

func TestGetAll(t *testing.T) {
	getPosts, err := get("posts", nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(getPosts))
}

func TestPost(t *testing.T) {
	postBody := map[string]string{
		"title": "foo",
		"body": "bar",
		"userId": "1",
	}

	createPost, err := post("posts", postBody)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(createPost))
}
func TestGetOne(t *testing.T) {
	getOne, err := get("posts", url.Values{
		"id": {strconv.Itoa(1)},
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(getOne))
}


