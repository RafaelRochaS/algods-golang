package operations

import (
	"algods/common"
	"algods/datastructures"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func HashTableMenu() {
	var input int
	for {
		common.PrintMenu(`0 - Return
1 - Test InsertRandom`)
		fmt.Scanf("%d", &input)
		switch input {
		case 0:
			return
		case 1:
			testInsertRandom()
		}
	}
}

func testInsertRandom() {
	ht := datastructures.NewHashTable()
	seed := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(seed)

	for {
		number := generator.Intn(1000)
		ht.Insert("key-"+strconv.Itoa(number), number)
	}
}
