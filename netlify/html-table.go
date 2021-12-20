package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/Jeffail/gabs"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

type people struct {
	Number string `json:"id"`
}

func main() {

	// make a sample HTTP GET request

	// read all response body
	data := Request("https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item?select=id,url,item_adult(value,score),item_description(value,score)&and=(id.eq.beec7a1a-bdc1-41dc-b7e1-5c15565c3efc)")
	// print `data` as a string
	fmt.Printf("%s\n", data)

	textBytes := []byte(data)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Number)

	value := gjson.Get(string(data), "item_description.0.value")
	println(value.String())

	/*
	   jsonParsed, err := gabs.ParseJSON(data)
	   if err != nil {
	     panic(err)
	   }
	   var blah string
	   var value string
	   var score string
	   blah = "<table>"
	   blah = blah + "<tr>"
	   value = jsonParsed.Path("id").Data().(string)
	   score = jsonParsed.Path("item_adult.score").Data().(string)
	   blah = blah + "<td>Item ID:</td><td>"+value+"</td>"
	   blah = blah + "<td>Item Score:</td><td>"+score+"</td>"
	   blah = blah + "</tr>"
	   blah = blah + "<tr>"
	   blah = blah + "<td>Score:</td><td>"+jsonParsed.Path("item_description.score").Data().(string)+"</td>"
	   value = jsonParsed.Path("item_description.value").Data().(string)
	   blah = blah + "<td>Item description:</td><td>"+value+"</td>"
	   value = jsonParsed.Path("item_description.value").Data().(string)
	   blah = blah + "<td>Item description:</td><td>"+value+"</td>"
	   blah = blah + "</tr>"
	   blah = blah + "<tr>"
	   blah = blah + "<td>Score:</td><td>"+jsonParsed.Path("item_description.score").Data().(string)+"</td>"
	   fmt.Printf( "%s\n", blah )
	*/
}

func Request(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYW5vbiIsImlhdCI6MTYwOTE1MTk4NiwiZXhwIjoxOTI0NzI3OTg2fQ.5Bjcho1pKyUaPfFoqHGQjOukeErw5k-nTeeFAVeNtVg")
	req.Header.Set("Accept", "application/vnd.pgrst.object+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
