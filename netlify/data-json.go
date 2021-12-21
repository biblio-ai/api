package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tidwall/gjson"
)

func json_handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// read all response body
	//Get the path parameter that was sent
	//uuid := request.PathParameters["uuid"]
	metadata_key := request.QueryStringParameters["metadata_key"]
	metadata_value := request.QueryStringParameters["metadata_value"]
        metadata_model := request.QueryStringParameters["metadata_model"]

        var data_id []byte
        if len(metadata_model) == 0 {
          data_id = Request("https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item_metadata_transpose?"+metadata_key+"=eq."+metadata_value+"&order=model.desc.nullslast&limit=1")
        } else {
          data_id = Request("https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item_metadata_transpose?"+metadata_key+"=eq."+metadata_value+"&and=(model.eq."+metadata_model+")&limit=1")
        }

	request_data_id := gjson.Get(string(data_id), "item_id").String()

	data := Request("https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item?select=id,url,item_brand(value,score,x,y,width,height),item_category(value,score),item_celebrity(value,score,position_height,position_left,position_top,position_width),item_color(black_and_white,accent_color,dominant_color_background,dominant_color_foreground,dominant_colors),item_description(value,score),item_face(position_height,position_left,position_top,position_width),item_landmark(value,score),item_object(value,score,x,y,width,height),item_tag(value,score),item_text(line,value,box),item_text_entity(value,match_text,text_type,text_sub_type,text_score),item_text_key_phrase(value),item_text_language(value,code,score),item_text_sentiment(score),item_metadata(metadata_key,metadata_value),item_log(section,value)&and=(id.eq." + request_data_id + ")")
	// print `data` as a string
	blah := string(data)
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{
					"Content-Type": "text/json",
				  	"Access-Control-Allow-Origin": "*",
					"Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept",
		},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              blah,
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	lambda.Start(json_handler)
	// make a sample HTTP GET request
}

func Request(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("apikey", os.Getenv("SUPABASE_API_KEY"))
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
