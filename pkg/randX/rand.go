package randX

import (
	"math/rand"
	"single-sentence/pkg/stringsX"
	"strings"
)

func GetRandParam(params []string) string {
	params = stringsX.RemoveDuplicates(params)
	allQueries := "abcdefghijkl"
	//fmt.Println(params)
	if len(params) == 0 {
		return string(allQueries[rand.Intn(len(allQueries))])
	}
	param := params[rand.Intn(len(params))]
	if !strings.Contains(allQueries, param) {
		return ""
	}
	return param
}
