package tag_api

import (
	"fmt"
	"reflect"
	"strings"
)

func (data *ApiData) AutoLoad() {

	// Load images
	data.LoadImages()
}

// Helper functions

func (data *ApiData) MakeSelect(dt interface{}) (sel string) {
	var tag, sql_tag string
	var ok bool
	var i int
	var fields []string

	st := reflect.TypeOf(dt)

	// Iterate through struct fields, gathering db tags
	for i = 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if tag, ok = field.Tag.Lookup("db"); ok {
			if tag != "" {
				// If sql tag specified, override db tag
				sql_tag, ok = field.Tag.Lookup("sql")
				if ok {
					if sql_tag == "-" {
						logging.Debug.Printf("Skipping field: %s [%s]\n", field.Name, field.Type)
						continue
					}
					tag = sql_tag + " AS " + tag
				}
				logging.Debug.Printf("Select %q for field: %s [%s]\n", tag, field.Name, field.Type)
				fields = append(fields, tag)
			}
		}
	}

	// Auto generate SELECT statement from tags
	sel = "SELECT " + strings.Join(fields, ", ")
	return
}

func (data *ApiData) MakeQuery(dt interface{}, query string, v ...interface{}) (finalq string) {

	// pad with a space unless query starts with a ','
	var pad string = " "
	if query[0] == ',' {
		pad = ""
	}

	finalq = data.MakeSelect(dt) + pad + fmt.Sprintf(query, v...)
	return
}
