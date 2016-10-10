package stats

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/module/uwsgi"
)

func getKey(name, prefix string) string {
	if prefix != "" {
		name = fmt.Sprintf("%s.%s", prefix, name)
	}
	return strings.ToLower(name)

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
				m = common.MapStrUnion(
					m,
					getFlatMapStr(
						val.Field(i).Index(j),
						getKey(fmt.Sprintf("%s.%d", typeOfT.Field(i).Name, j), prefix)))
			}
		case reflect.Struct:
			for j := 0; j < val.NumField(); j++ {
				m = common.MapStrUnion(m, getFlatMapStr(val.Field(j), typeOfT.Field(i).Name))
			}
		default:
			fmt.Println("type switch not found", val.Field(i).Kind())

		}

	}
	return m
}

func getMapStr(stat *uwsgi.Stat) common.MapStr {
	val := reflect.ValueOf(stat).Elem()
	m := common.MapStr{}
	typeOfT := val.Type()
	for i := 0; i < val.NumField(); i++ {
		m[typeOfT.Field(i).Name] = val.Field(i)
	}
	return m
}

// Map data to MapStr
func eventMapping(stat *uwsgi.Stat) common.MapStr {
	//var event common.MapStr

	//val := reflect.ValueOf(stat).Elem()
	//event = getFlatMapStr(val, "")
	//event := structs.Map(stat)
	event := getMapStr(stat)

	return event
}
