package randX

import (
	"fmt"
	"testing"
)

var testList = []string{"1", "2", "3", "4", "5"}
var testList2 = make([]string, 0)

func TestGetRandParam(t *testing.T) {
	fmt.Println(len(testList2))
	for i := 0; i < 10; i++ {
		fmt.Println(GetRandParam(testList2))
	}

}
