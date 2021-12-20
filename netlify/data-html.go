package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
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
          data_id = Request("https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item_metadata_transpose?"+metadata_key+"=eq."+metadata_value+"&and=(model.eq."+metadata_model+"&limit=1)")
        }

        //	data_id := Request( "https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item_metadata?metadata_key=eq."+metadata_key+"&metadata_value=eq."+metadata_value )
        request_data_id :=gjson.Get(string(data_id), "item_id").String()

        data := Request("https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item?select=id,url,item_adult(value,score),item_brand(value,score,x,y,width,height),item_category(value,score),item_celebrity(value,score,position_height,position_left,position_top,position_width),item_color(black_and_white,accent_color,dominant_color_background,dominant_color_foreground,dominant_colors),item_description(value,score),item_face(gender,age,position_height,position_left,position_top,position_width),item_landmark(value,score),item_object(value,score,x,y,width,height),item_racy(value,score),item_tag(value,score),item_text(line,value,box),item_text_entity(value,match_text,text_type,text_sub_type,text_score),item_text_key_phrase(value),item_text_language(value,code,score),item_text_sentiment(score),item_metadata(metadata_key,metadata_value),item_log(section,value)&and=(id.eq." + request_data_id + ")")
        // print `data` as a string
        //blah := string(data)
        // value := gjson.Get(string(data), "item_description.0.value")
        var blah string
        blah = "<html><head><link rel='stylesheet' href='https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css' integrity='sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T' crossorigin='anonymous'></head><body style='padding:15px;'><table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead>"
        blah = blah + "<tr><th colspan='2'>API data:</th></tr>"
        blah = blah + "</thead><tbody>"
        blah = blah + "<tr><td style='width:20%;'>Key:</td><td>" + metadata_key + "</td></tr>"
        blah = blah + "<tr><td>Value:</td><td>" + metadata_value + "</td></tr>"
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Description:</th><tr>"
        blah = blah + "</thead><tbody>"
        blah = blah + "<tr>"
        blah = blah + "<td style='width:20%;'>Description:</td><td>" + gjson.Get(string(data), "item_description.0.value").String() + "</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah_percentage, _ := strconv.ParseFloat(gjson.Get(string(data), "item_description.0.score").String(), 8)
        blah = blah + "<td>Score:</td><td>" + fmt.Sprintf("%.2f", 100*blah_percentage) + "%</td>"
        blah = blah + "</tr>"
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Tags:</th><tr>"
        blah = blah + "</thead><tbody>"
        type Tdata struct {
          Value string
          Score string
        }
        var Tdatas []Tdata

        result := gjson.Get(string(data), "item_tag")
        json.Unmarshal([]byte(result.Raw), &Tdatas)
        for _, name := range Tdatas {
          blah = blah + "<tr>"
          blah = blah + "<td style='width:20%;'>" + name.Value + "</td>"
          blah_score, _ := strconv.ParseFloat(name.Score, 32)
          blah = blah + "<td>Score: " + fmt.Sprintf("%.2f", 100*blah_score) + "%</td>"
          blah = blah + "</tr>"
        }
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Objects:</th><tr>"
        blah = blah + "</thead><tbody>"
        type Odata struct {
          Value  string
          Score  string
          X      string
          Y      string
          Width  string
          Height string
        }
        var Odatas []Odata

        result_obj := gjson.Get(string(data), "item_object")
        json.Unmarshal([]byte(result_obj.Raw), &Odatas)
        for _, name := range Odatas {
          blah = blah + "<tr>"
          blah = blah + "<td style='width:20%;'>" + name.Value + "</td>"
          blah_score, _ := strconv.ParseFloat(name.Score, 32)
          blah = blah + "<td>Score: " + fmt.Sprintf("%.2f", 100*blah_score) + "% Location: (" + name.X + " " + name.Y + " " + name.Width + " " + name.Height + ")</td>"
          blah = blah + "</tr>"
        }
        blah = blah + "<tr>"
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Text:</th><tr>"
        blah = blah + "</thead><tbody>"
        result_text := gjson.Get(string(data), "item_text.#.value")
        blah = blah + "<tr>"
        blah = blah + "<td colspan='2'>"
        for _, name := range result_text.Array() {
          blah = blah + name.String() + "<br />"
        }
        blah = blah + "</td>"
        blah = blah + "</tr>"
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Key phrases:</th><tr>"
        blah = blah + "</thead><tbody>"
        result_text_key := gjson.Get(string(data), "item_text_key_phrase.#.value")
        for _, name := range result_text_key.Array() {
          blah = blah + "<tr>"
          blah = blah + "<td colspan='2'>" + name.String() + "</td>"
          blah = blah + "</tr>"
        }
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Faces:</th><tr>"
        blah = blah + "</thead><tbody>"
        type Fdata struct {
          Position_height string
          Position_left   string
          Position_top    string
          Position_width  string
        }
        var Fdatas []Fdata

        result_face := gjson.Get(string(data), "item_face")
        json.Unmarshal([]byte(result_face.Raw), &Fdatas)
        for _, name := range Fdatas {
          blah = blah + "<tr>"
          blah = blah + "<td style='width:20%;'>Location: </td>"
          blah = blah + "<td> " + name.Position_height + " " + name.Position_left + " " + name.Position_top + " " + name.Position_width + "</td>"
          blah = blah + "</tr>"
        }
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Entities:</th><tr>"
        blah = blah + "</thead><tbody>"
        type Edata struct {
          Value         string
          Match_text    string
          Text_type     string
          Text_sub_type string
          Text_score    string
        }
        var Edatas []Edata

        result_ent := gjson.Get(string(data), "item_text_entity")
        json.Unmarshal([]byte(result_ent.Raw), &Edatas)
        for _, name := range Edatas {
          blah = blah + "<tr>"
          blah = blah + "<td style='width:20%;'>" + name.Value + "</td>"
          blah_score, _ := strconv.ParseFloat(name.Text_score, 32)
          blah = blah + "<td>Score: " + fmt.Sprintf("%.2f", 100*blah_score) + "% Match text: " + name.Match_text + "  Type: " + name.Text_type + " Sub-Type: " + name.Text_sub_type + "</td>"
          blah = blah + "</tr>"
        }
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Metadata:</th><tr>"
        blah = blah + "</thead><tbody>"

        type Mdata struct {
          Metadata_key   string
          Metadata_value string
        }
        var Mdatas []Mdata

        result_metadata := gjson.Get(string(data), "item_metadata")
        json.Unmarshal([]byte(result_metadata.Raw), &Mdatas)
        for _, name := range Mdatas {
          blah = blah + "<tr>"
          blah = blah + "<td style='width:20%;'>" + name.Metadata_key + ":</td>"
          blah = blah + "<td>" + name.Metadata_value + "</td>"
          blah = blah + "</tr>"
        }
        blah = blah + "</tbody></table>"

        blah = blah + "<br />"

        blah = blah + "<table class='table table-striped table-bordered table-sm table'>"
        blah = blah + "<thead><tr><th colspan='2'>Logs:</th><tr>"
        blah = blah + "</thead><tbody>"

        type Ldata struct {
          Section string
          Value   string
        }
        var Ldatas []Ldata

        result_log := gjson.Get(string(data), "item_log")
        json.Unmarshal([]byte(result_log.Raw), &Ldatas)
        for _, name := range Ldatas {
          blah = blah + "<tr>"
          blah = blah + "<td style='width:20%;'>" + name.Section + "</td>"
          blah = blah + "<td>" + name.Value + "</td>"
          blah = blah + "</tr>"
        }

        blah = blah + "</tbody></table>"
        // Iterating address objects
        /*
        for key, child := range jsonParsed.Search("employees", "address").ChildrenMap() {
          fmt.Printf("Key=>%v, Value=>%v\n", key, child.Data().(string))
        }
        */

        return &events.APIGatewayProxyResponse{
          StatusCode:        200,
          Headers:           map[string]string{
            "Content-Type": "text/html",
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept",
          },
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
