# Biblio.AI API

The Biblio.ai API is a collection of serverless functions - currently running on Netlify, _though can be ported easily to AWS Lambda etc.._ to provide the AI enriched metadata for a library collection item, based on simple API/URL parameters

## Example

Assuming the image `cor-green_to_swinburne_1920.jpg` has already been imported and AI enriched, we can use the below values to query the API for the metadta

* `metadata_key` = `metadata_identifier_file_id`
* `metadata_value` = `cor-green_to_swinburne_1920.jpg`

`metadata_key` can be any of the values imported into Biblio.AI for that image - from the OAI feed - such and handle, id etc..

## Format

Currently HTML and JSON are supported outputs - though MARCXML and DC are planned

For HTML - use `data-html` the in url
For JSON use `data-json` in the url

## HTML

Example URL query

```
https://api.biblio.ai/.netlify/functions/data-html?metadata_key=metadata_identifier_file_id&metadata_value=cor-green_to_swinburne_1920.jpg
```
HTML output

![image](https://user-images.githubusercontent.com/249520/137823169-1b313a0e-2ad2-4f87-9e88-9a29ad0c04e3.png)


## JSON

Example URL query
```
https://api.biblio.ai/.netlify/functions/data-json?metadata_key=metadata_identifier_file_id&metadata_value=cor-green_to_swinburne_1920.jpg
```

JSON output
```
{
  "id": "2232bcca-e3f3-4f30-8499-5e8f9b6ef9e9",
  "url": "https://commons.swinburne.edu.au/file/e8dab624-fa0c-47d8-a88b-066a6d8f9314/1/cor-green_to_swinburne_1920.jpg",
  "item_brand": [],
  "item_category": [
    {
      "value": "text_",
      "score": "0.95703125"
    }
  ],
  "item_celebrity": [],
  "item_color": [],
  "item_description": [
    {
      "value": "text, letter",
      "score": "0.9999932050704956"
    }
  ],
  "item_face": [],
  "item_landmark": [],
  "item_object": [],
  "item_tag": [
    {
      "value": "text",
      "score": "0.9998282790184021"
    },
    {
      "value": "letter",
      "score": "0.8675053119659424"
    },
    {
      "value": "screenshot",
      "score": "0.7274935245513916"
    },
    {
      "value": "document",
      "score": "0.2567472457885742"
    }
  ],
  "item_text": [
    {
      "line": 0,
      "value": "15th. June, 1920.",
      "box": "[1721 164 2230 164 2230 211 1721 211]"
    },
    {
      "line": 1,
      "value": "Honorable G. Swinburne,",
      "box": "[184 478 876 476 876 525 184 526]"
    },
    {
      "line": 2,
      "value": "\"Shenton\"",
      "box": "[304 519 590 521 589 570 303 569]"
    },
    {
      "line": 3,
      "value": "hinkora hosd,",
      "box": "[397 579 798 575 798 623 397 626]"
    },
    {
      "line": 4,
      "value": "Hawthorn.",
      "box": "[897 572 1182 573 1182 619 897 618]"
    },
    {
      "line": 5,
      "value": "Dear Mr. Swinburne,",
      "box": "[190 676 760 678 759 724 190 721]"
    },
    {
      "line": 6,
      "value": "I have been approached on one or two occasions, and just",
      "box": "[486 777 2146 763 2146 811 486 825]"
    },
    {
      "line": 7,
      "value": "recently by the Engineer in charge of the Government Cool Stores",
      "box": "[191 826 2111 812 2112 863 191 877]"
    },
    {
      "line": 8,
      "value": "at Victoria Docks, in regard to the training of refrigerating",
      "box": "[196 872 2031 862 2031 915 196 925]"
    },
    {
      "line": 9,
      "value": "engineers.",
      "box": "[196 930 490 928 490 971 196 972]"
    },
    {
      "line": 10,
      "value": "One of the suggestions is to include in our Day Engineer-",
      "box": "[565 921 2303 909 2303 961 565 973]"
    },
    {
      "line": 11,
      "value": "ing classes a course of work in \"Regrigerating Engineering\".",
      "box": "[191 974 2003 962 2003 1013 191 1025]"
    },
    {
      "line": 12,
      "value": "Such",
      "box": "[2046 962 2183 965 2182 1009 2045 1006]"
    },
    {
      "line": 13,
      "value": "a course would run parallel with our other diploma courses, and would",
      "box": "[188 1025 2269 1007 2270 1059 188 1077]"
    },
    {
      "line": 14,
      "value": "be a means of training regrigerating engineers who will be much needed",
      "box": "[194 1075 2304 1057 2304 1108 194 1126]"
    },
    {
      "line": 15,
      "value": "in Victoria in the future,",
      "box": "[198 1122 995 1120 995 1171 198 1173]"
    },
    {
      "line": 16,
      "value": "The Institution of Refrigerating",
      "box": "[1079 1113 2063 1112 2063 1163 1079 1165]"
    },
    {
      "line": 17,
      "value": "Engineers are interested in the matter and are willing to give all the",
      "box": "[198 1173 2299 1156 2299 1205 199 1222]"
    },
    {
      "line": 18,
      "value": "help they can, and to provide a Scholarship as well.",
      "box": "[190 1224 1753 1212 1754 1261 190 1273]"
    },
    {
      "line": 19,
      "value": "We should require a small experimental plant for testing",
      "box": "[514 1320 2216 1308 2216 1356 514 1368]"
    },
    {
      "line": 20,
      "value": "and research work, and as no College in the Melbourne district has",
      "box": "[193 1373 2173 1357 2174 1407 193 1423]"
    },
    {
      "line": 21,
      "value": "such a plant the research work, which could be done, would I think be",
      "box": "[194 1425 2275 1406 2276 1456 194 1475]"
    },
    {
      "line": 22,
      "value": "valuable.",
      "box": "[196 1473 483 1474 483 1521 196 1520]"
    },
    {
      "line": 23,
      "value": "I think Messrs. Werner would help us, and might be glad",
      "box": "[579 1472 2251 1459 2251 1509 580 1522]"
    },
    {
      "line": 24,
      "value": "to have a plant handy in order to carry out tests, and I am getting",
      "box": "[203 1524 2218 1511 2219 1560 203 1574]"
    },
    {
      "line": 25,
      "value": "in touch with them.",
      "box": "[202 1571 778 1571 778 1620 202 1620]"
    },
    {
      "line": 26,
      "value": "You mentioned some time ago about having had given to you",
      "box": "[520 1669 2244 1660 2244 1711 520 1720]"
    },
    {
      "line": 27,
      "value": "some $500 for the College, and the thought occurred to me that you",
      "box": "[197 1722 2192 1711 2192 1762 197 1773]"
    },
    {
      "line": 28,
      "value": "might be willing to have it expended to help forward the scheme.",
      "box": "[193 1769 2128 1759 2128 1811 194 1822]"
    },
    {
      "line": 29,
      "value": "Perhaps Mr. Kerr's ground could be used as the cooling rooms would",
      "box": "[199 1823 2177 1810 2178 1860 199 1873]"
    },
    {
      "line": 30,
      "value": "probably be built of wood, as is usually the case in Victoria.",
      "box": "[198 1872 2065 1859 2065 1909 198 1922]"
    },
    {
      "line": 31,
      "value": "If you think favorably of the scheme I thought of getting",
      "box": "[523 1970 2246 1961 2246 2012 523 2021]"
    },
    {
      "line": 32,
      "value": "some refrigerating engineers together and talking the matter over",
      "box": "[202 2025 2159 2010 2159 2060 202 2076]"
    },
    {
      "line": 33,
      "value": "with them, and perhaps some of the refrigerating firms might be will-",
      "box": "[200 2072 2276 2057 2276 2109 200 2124]"
    },
    {
      "line": 34,
      "value": "ing to help us financially,",
      "box": "[203 2123 1021 2119 1021 2171 203 2175]"
    },
    {
      "line": 35,
      "value": "I think there is a large field for research and experimental",
      "box": "[525 2218 2333 2207 2333 2256 526 2267]"
    },
    {
      "line": 36,
      "value": "work in connection with refrigeration in Victoria as so many districts",
      "box": "[205 2273 2306 2258 2306 2306 206 2321]"
    },
    {
      "line": 37,
      "value": "are now having cooling stores erected.",
      "box": "[211 2324 1316 2311 1316 2358 211 2371]"
    },
    {
      "line": 38,
      "value": "With kind regards,",
      "box": "[711 2414 1252 2414 1252 2464 711 2464]"
    },
    {
      "line": 39,
      "value": "I all,",
      "box": "[992 2485 1141 2494 1138 2537 989 2528]"
    },
    {
      "line": 40,
      "value": "Faithfully yours,",
      "box": "[1048 2546 1559 2551 1558 2601 1047 2596]"
    }
  ],
  "item_text_entity": [
    {
      "value": "15th",
      "match_text": "15th",
      "text_type": "Quantity",
      "text_sub_type": "Ordinal",
      "text_score": "0.8"
    },
    {
      "value": "June, 1920",
      "match_text": "June, 1920",
      "text_type": "DateTime",
      "text_sub_type": "DateRange",
      "text_score": "0.8"
    },
    {
      "value": "The Honourable",
      "match_text": "Honorable",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "G. Swinburne",
      "match_text": "G. Swinburne",
      "text_type": "Person",
      "text_sub_type": "",
      "text_score": "0.7"
    },
    {
      "value": "Swinburne University of Technology",
      "match_text": "Swinburne",
      "text_type": "Organization",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Swinburne University of Technology",
      "match_text": "Swinburne",
      "text_type": "Organization",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Hawthorn.\nDear",
      "match_text": "Hawthorn.\nDear",
      "text_type": "Organization",
      "text_sub_type": "",
      "text_score": "0.95"
    },
    {
      "value": "Hawthorn, Victoria",
      "match_text": "Hawthorn",
      "text_type": "Location",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Swinburne",
      "match_text": "Swinburne",
      "text_type": "Person",
      "text_sub_type": "",
      "text_score": "1"
    },
    {
      "value": "one",
      "match_text": "one",
      "text_type": "Quantity",
      "text_sub_type": "Number",
      "text_score": "0.8"
    },
    {
      "value": "two",
      "match_text": "two",
      "text_type": "Quantity",
      "text_sub_type": "Number",
      "text_score": "0.8"
    },
    {
      "value": "recently",
      "match_text": "recently",
      "text_type": "DateTime",
      "text_sub_type": "",
      "text_score": "0.8"
    },
    {
      "value": "Cool (aesthetic)",
      "match_text": "Cool",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Victoria Docks",
      "match_text": "Victoria Docks",
      "text_type": "Location",
      "text_sub_type": "GPE",
      "text_score": "0.7"
    },
    {
      "value": "One",
      "match_text": "One",
      "text_type": "Quantity",
      "text_sub_type": "Number",
      "text_score": "0.8"
    },
    {
      "value": "Our Day",
      "match_text": "our Day",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Engineering",
      "match_text": "Engineering",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Engineering",
      "match_text": "Engineers",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Engineering",
      "match_text": "Engineer",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Victoria (Australia)",
      "match_text": "Victoria",
      "text_type": "Location",
      "text_sub_type": "GPE",
      "text_score": "0.99"
    },
    {
      "value": "Victoria (Australia)",
      "match_text": "Victoria",
      "text_type": "Location",
      "text_sub_type": "GPE",
      "text_score": "0.96"
    },
    {
      "value": "Victoria (Australia)",
      "match_text": "Victoria",
      "text_type": "Location",
      "text_sub_type": "GPE",
      "text_score": "0.88"
    },
    {
      "value": "The Institution of Refrigerating\nEngineers",
      "match_text": "The Institution of Refrigerating\nEngineers",
      "text_type": "Organization",
      "text_sub_type": "",
      "text_score": "0.96"
    },
    {
      "value": "Institution",
      "match_text": "Institution",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Scholarship",
      "match_text": "Scholarship",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "College",
      "match_text": "College",
      "text_type": "Organization",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Melbourne",
      "match_text": "Melbourne",
      "text_type": "Location",
      "text_sub_type": "GPE",
      "text_score": "1"
    },
    {
      "value": "Werner",
      "match_text": "Werner",
      "text_type": "Person",
      "text_sub_type": "",
      "text_score": "1"
    },
    {
      "value": "$500",
      "match_text": "$500",
      "text_type": "Quantity",
      "text_sub_type": "Currency",
      "text_score": "0.8"
    },
    {
      "value": "Working Men's College, Melbourne",
      "match_text": "the College",
      "text_type": "Organization",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Metropolitan Railway",
      "match_text": "Mr.",
      "text_type": "Location",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Metropolitan Railway",
      "match_text": "Mr.",
      "text_type": "Location",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "Kerr",
      "match_text": "Kerr",
      "text_type": "Person",
      "text_sub_type": "",
      "text_score": "0.98"
    },
    {
      "value": "Wood",
      "match_text": "wood",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    },
    {
      "value": "now",
      "match_text": "now",
      "text_type": "DateTime",
      "text_sub_type": "",
      "text_score": "0.8"
    },
    {
      "value": "Faithfully Yours",
      "match_text": "Faithfully yours",
      "text_type": "Other",
      "text_sub_type": "",
      "text_score": "0"
    }
  ],
  "item_text_key_phrase": [
    {
      "value": "refrigerating engineers"
    },
    {
      "value": "research work"
    },
    {
      "value": "course of work"
    },
    {
      "value": "training of refrigerating"
    },
    {
      "value": "Victoria Docks"
    },
    {
      "value": "Institution of Refrigerating"
    },
    {
      "value": "refrigerating firms"
    },
    {
      "value": "means of training regrigerating engineers"
    },
    {
      "value": "small experimental plant"
    },
    {
      "value": "plant handy"
    },
    {
      "value": "College"
    },
    {
      "value": "matter"
    },
    {
      "value": "cooling stores"
    },
    {
      "value": "Swinburne"
    },
    {
      "value": "thought"
    },
    {
      "value": "scheme"
    },
    {
      "value": "Regrigerating Engineering"
    },
    {
      "value": "Day Engineer"
    },
    {
      "value": "ing classes"
    },
    {
      "value": "cooling rooms"
    },
    {
      "value": "Government Cool Stores"
    },
    {
      "value": "charge"
    },
    {
      "value": "refrigeration"
    },
    {
      "value": "order"
    },
    {
      "value": "testing"
    },
    {
      "value": "large field"
    },
    {
      "value": "connection"
    },
    {
      "value": "diploma courses"
    },
    {
      "value": "tests"
    },
    {
      "value": "Melbourne district"
    },
    {
      "value": "case"
    },
    {
      "value": "districts"
    },
    {
      "value": "Kerr's ground"
    },
    {
      "value": "kind regards"
    },
    {
      "value": "hinkora hosd"
    },
    {
      "value": "suggestions"
    },
    {
      "value": "wood"
    },
    {
      "value": "future"
    },
    {
      "value": "Shenton"
    },
    {
      "value": "Hawthorn"
    },
    {
      "value": "occasions"
    },
    {
      "value": "Scholarship"
    },
    {
      "value": "Werner"
    },
    {
      "value": "touch"
    },
    {
      "value": "time"
    }
  ],
  "item_text_language": [
    {
      "value": "English",
      "code": "en",
      "score": "1"
    }
  ],
  "item_text_sentiment": [],
  "item_metadata": [
    {
      "metadata_key": "url",
      "metadata_value": "https://commons.swinburne.edu.au/file/e8dab624-fa0c-47d8-a88b-066a6d8f9314/1/cor-green_to_swinburne_1920.jpg"
    },
    {
      "metadata_key": "name",
      "metadata_value": "swin-trovetest"
    },
    {
      "metadata_key": "url",
      "metadata_value": "https://commons.swinburne.edu.au/file/e8dab624-fa0c-47d8-a88b-066a6d8f9314/1/cor-green_to_swinburne_1920.jpg"
    },
    {
      "metadata_key": "header_identifier",
      "metadata_value": "oai:commons.swinburne.edu.au:e8dab624-fa0c-47d8-a88b-066a6d8f9314/1"
    },
    {
      "metadata_key": "date_latest",
      "metadata_value": "2021-09-09T02:48:12Z"
    },
    {
      "metadata_key": "metadata_identifier",
      "metadata_value": "e8dab624-fa0c-47d8-a88b-066a6d8f9314"
    },
    {
      "metadata_key": "metadata_identifier_file_id",
      "metadata_value": "cor-green_to_swinburne_1920.jpg"
    }
  ],
  "item_log": []
}
```
