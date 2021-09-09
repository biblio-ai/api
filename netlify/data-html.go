package main

import (
  	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"log"
	"os"
	"io/ioutil"
        "github.com/tidwall/gjson"
        "encoding/json"
)
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// read all response body
      	//Get the path parameter that was sent
	//uuid := request.PathParameters["uuid"]
        metadata_key := request.QueryStringParameters["metadata_key"]
        metadata_value := request.QueryStringParameters["metadata_value"]

	data_id := Request( "https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item_metadata?metadata_key=eq."+metadata_key+"&metadata_value=eq."+metadata_value )
        request_data_id :=gjson.Get(string(data_id), "item_id").String()

	data := Request( "https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item?select=id,url,item_adult(value,score),item_brand(value,score,x,y,width,height),item_category(value,score),item_celebrity(value,score,position_height,position_left,position_top,position_width),item_color(black_and_white,accent_color,dominant_color_background,dominant_color_foreground,dominant_colors),item_description(value,score),item_face(gender,age,position_height,position_left,position_top,position_width),item_landmark(value,score),item_object(value,score,x,y,width,height),item_racy(value,score),item_tag(value,score),item_text(line,value,box),item_text_entity(value,match_text,text_type,text_sub_type,text_score),item_text_key_phrase(value),item_text_language(value,code,score),item_text_sentiment(score),item_metadata(metadata_key,metadata_value)&and=(id.eq."+request_data_id+")")
	// print `data` as a string
        //blah := string(data)
       // value := gjson.Get(string(data), "item_description.0.value")
        var blah string
        blah = "<table>"
        blah = blah + "<tr><td>Key:</td><td>"+metadata_key+"</td></tr>"
        blah = blah + "<tr><td>Value:</td><td>"+metadata_value+"</td></tr>"
        blah = blah + "<tr><td>Origin tem ID:</td><td>"+gjson.Get(string(data_id), "item_id").String()+"</td></tr>"
        blah = blah + "<tr>" 
        blah = blah + "<td>Item ID:</td><td>"+gjson.Get(string(data), "id").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td>Item URL:</td><td>"+gjson.Get(string(data), "url").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item description:</b></td><td></td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td>Description:</td><td>"+gjson.Get(string(data), "item_description.0.value").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td>Score:</td><td>"+gjson.Get(string(data), "item_description.0.score").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item tag:</b></td><td></td>"
        blah = blah + "</tr>"
        result := gjson.Get(string(data), "item_tag.#.value")
for _, name := range result.Array() {
                blah = blah + "<tr>"
                blah = blah + "<td colspan='2'>"+name.String()+"</td>"
                blah = blah + "</tr>"
}
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item object:</b></td><td></td>"
        blah = blah + "</tr>"
        type Odata struct {
          Value     string
          Score string
          x string
          y string
          width string
          height string
        }
        var Odatas []Odata

        result_obj := gjson.Get(string(data), "item_object.#.value")
        json.Unmarshal([]byte(result_obj.Raw), &Odatas)
        for _, name := range Odatas {
                blah = blah + "<tr>"
                blah = blah + "<td>"+name.Value+"</td>"
                blah = blah + "<td>"+name.Score+" Location: ("+name.x+" "+name.y+" "+name.width+" "+name.height+")</td>"
                blah = blah + "</tr>"
}
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item text:</b></td><td></td>"
        blah = blah + "</tr>"
        result_text := gjson.Get(string(data), "item_text.#.value")
for _, name := range result_text.Array() {
                blah = blah + "<tr>"
                blah = blah + "<td colspan='2'>"+name.String()+"</td>"
                blah = blah + "</tr>"
}
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item text key phrases:</b></td><td></td>"
        blah = blah + "</tr>"
        result_text_key := gjson.Get(string(data), "item_text_key_phrase.#.value")
for _, name := range result_text_key.Array() {
                blah = blah + "<tr>"
                blah = blah + "<td colspan='2'>"+name.String()+"</td>"
                blah = blah + "</tr>"
}
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item metadata:</b></td><td></td>"
        blah = blah + "</tr>"

        type Mdata struct {
          Metadata_key     string
          Metadata_value string
        }
        var Mdatas []Mdata

        result_metadata := gjson.Get(string(data), "item_metadata")
        json.Unmarshal([]byte(result_metadata.Raw), &Mdatas)
        for _, name := range Mdatas {
          blah = blah + "<tr>"
          blah = blah + "<td>"+name.Metadata_key+":</td>"
          blah = blah + "<td>"+name.Metadata_value+"</td>"
          blah = blah + "</tr>"
        }
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item faces:</b></td><td></td>"
        blah = blah + "</tr>"
        type Fdata struct {
          Gender     string
          Age string
          Position_height string
          Position_left string
          Position_top string
          Position_width string
        }
        var Fdatas []Fdata

        result_face := gjson.Get(string(data), "item_face")
        json.Unmarshal([]byte(result_face.Raw), &Fdatas)
        for _, name := range Fdatas {
          blah = blah + "<tr>"
          blah = blah + "<td>Gender: "+name.Gender+"</td>"
          blah = blah + "<td>Age: "+name.Age+" Location:("+name.Position_height+" "+name.Position_left+" "+name.Position_top+" "+name.Position_width+")</td>"
          blah = blah + "</tr>"
        }

        blah = blah + "</table>"
        // Iterating address objects
        /*
        for key, child := range jsonParsed.Search("employees", "address").ChildrenMap() {
          fmt.Printf("Key=>%v, Value=>%v\n", key, child.Data().(string))
        }
        */

        return &events.APIGatewayProxyResponse{
          StatusCode:        200,
          Headers:           map[string]string{"Content-Type": "text/html"},
          MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
          Body:              string(blah),
          IsBase64Encoded:   false,
        }, nil
      }

      func main() {

        lambda.Start(handler)
        // make a sample HTTP GET request


      }


      func Request(url string) []byte {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
          log.Fatal(err)
        }
        req.Header.Set("apikey",  os.Getenv("SUPABASE_API_KEY"))
        req.Header.Set("Accept",  "application/vnd.pgrst.object+json")
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

