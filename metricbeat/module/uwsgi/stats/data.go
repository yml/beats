package stats

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/module/uwsgi"
)

func getKey(name, prefix string) string {
	if prefix != "" {
		name = fmt.Sprintf("%s.%s", prefix, name)
	}
	return name

}

func getFlatMapStr(val reflect.Value, prefix string) common.MapStr {
	m := common.MapStr{}
	typeOfT := val.Type()
	for i := 0; i < val.NumField(); i++ {

		switch val.Field(i).Kind() {
		case reflect.Int:
			m[getKey(typeOfT.Field(i).Name, prefix)] = val.Field(i).Int()
		case reflect.String:
			m[getKey(typeOfT.Field(i).Name, prefix)] = val.Field(i).String()
		case reflect.Slice:
			for j := 0; j < val.Field(i).Len(); j++ {
				m = common.MapStrUnion(m, getFlatMapStr(val.Field(i).Index(j), fmt.Sprintf("%s.%d", typeOfT.Field(j).Name, j)))
				spew.Dump(val.Field(j))
			}
		case reflect.Struct:
			spew.Dump(val.Field(i))
			for j := 0; j < val.NumField(); j++ {
				m = common.MapStrUnion(m, getFlatMapStr(val.Field(j), typeOfT.Field(j).Name))
				spew.Dump(val.Field(j))
			}

		default:
			fmt.Println("type switch not found", val.Field(i).Kind())

		}

	}
	return m
}

// Map data to MapStr
func eventMapping(stat *uwsgi.Stat) common.MapStr {
	var event common.MapStr

	val := reflect.ValueOf(stat).Elem()
	event = getFlatMapStr(val, "")

	spew.Dump(event)
	return event
}
