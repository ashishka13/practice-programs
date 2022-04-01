package repeatprogs

import (
	"encoding/json"
	"fmt"
	"log"
)

func BlacklistAttribute(jsonMap map[string]interface{}, black string) {

	for key1, val1 := range jsonMap {
		if key1 == black {
			delete(jsonMap, black)
		}
		if val2, ok := val1.(map[string]interface{}); ok {
			BlacklistAttribute(val2, black)
		}
	}
}

func main() {

	jsonstr := `{
		"glossary": {
			"title": "example glossary",
			"GlossDiv": {
				"title": "S",
				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": ["GML", "XML"]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}`

	mymap := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonstr), &mymap)
	log.Println(err)
	BlacklistAttribute(mymap, "GlossList")

	fmt.Print(mymap)
}
