package _map_test

import (
	"fmt"
	"testing"

	_map "github.com/mstmdev/puzzle-go/map"
)

func TestMap(t *testing.T) {
	m := make(map[any]any)
	hm := _map.NewHashMap(m)
	m["id"] = 1001
	m["lang"] = "go"
	m["type"] = "map"
	m["is_valid"] = true
	for i := 0; i < 100; i++ {
		m[fmt.Sprintf("random_%d", i)] = 100 + i
		hm.ShowStat()
		if hm.OldBucketCount() > 0 {
			oldBuckets := hm.OldBuckets()
			fmt.Println(oldBuckets)

			allData := hm.AllData()
			fmt.Println(allData)

			goAllData := hm.GoMap().AllData()
			fmt.Println(goAllData)

			if len(allData) != len(goAllData) {
				t.Errorf("AllData expect get data length %d, but get %d", len(goAllData), len(allData))
			}
		}
	}

	h := hm.HMap()
	fmt.Println(h)

	buckets := hm.Buckets()
	fmt.Println(buckets)

	allData := hm.AllData()
	fmt.Println(allData)

	goAllData := hm.GoMap().AllData()
	fmt.Println(goAllData)

	overflowCount := hm.OverflowCount()
	fmt.Println(overflowCount)

	oldBuckets := hm.OldBuckets()
	fmt.Println(oldBuckets)
}
