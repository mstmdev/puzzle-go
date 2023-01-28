package _map

import (
	"fmt"
	"reflect"
	"unsafe"
)

type HashMap struct {
	m map[any]any
}

func NewHashMap(m map[any]any) *HashMap {
	return &HashMap{m: m}
}

func (hm *HashMap) GoMap() GoMap {
	return hm.m
}

func (hm *HashMap) HMap() hmap {
	return **(**hmap)(unsafe.Pointer(&hm.m))
}

func (hm *HashMap) BucketCount() int {
	return 1 << hm.B()
}

func (hm *HashMap) B() uint8 {
	return hm.HMap().B
}

func (hm *HashMap) Len() int {
	return hm.HMap().count
}

func (hm *HashMap) Flags() uint8 {
	return hm.HMap().flags
}

func (hm *HashMap) HashSeed() uint32 {
	return hm.HMap().hash0
}

func (hm *HashMap) Bucket(bucketIndex uintptr) bmap {
	return *(*bmap)(unsafe.Pointer(uintptr(hm.HMap().buckets) + reflect.TypeOf(bmap{}).Size()*bucketIndex))
}

func (hm *HashMap) FirstBucket() bmap {
	return *(*bmap)(hm.HMap().buckets)
}

func (hm *HashMap) Buckets() []bmap {
	var buckets []bmap
	count := hm.BucketCount()
	for i := 0; i < count; i++ {
		buckets = append(buckets, hm.Bucket(uintptr(i)))
	}
	return buckets
}

func (hm *HashMap) OverflowCount() int {
	count := 0
	buckets := hm.Buckets()
	for _, bucket := range buckets {
		count += hm.getOverflowCount(bucket.overflow)
	}
	return count
}

func (hm *HashMap) getOverflowCount(overflow uintptr) int {
	if overflow == 0 {
		return 0
	}
	count := 1
	bucket := hm.getOverflow(overflow)
	if bucket.overflow > 0 {
		count += hm.getOverflowCount(bucket.overflow)
	}
	return count
}

func (hm *HashMap) dataList(bucket bmap) []Data {
	if bucket.tophash[0] < minTopHash {
		return nil
	}
	var dataList []Data
	for i, bit := range bucket.tophash {
		if bit >= minTopHash {
			if bucket.keys[i] != nil && bucket.elems[i] != nil {
				dataList = append(dataList, newData(bucket.keys[i], bucket.elems[i]))
			}
		}
	}
	if bucket.overflow > 0 {
		overflow := hm.getOverflow(bucket.overflow)
		dataList = append(dataList, hm.dataList(overflow)...)
	}
	return dataList
}

func (hm *HashMap) getOverflow(overflow uintptr) bmap {
	return *(*bmap)(unsafe.Pointer(overflow))
}

func (hm *HashMap) AllData() []Data {
	var allData []Data
	buckets := hm.Buckets()
	for _, bucket := range buckets {
		allData = append(allData, hm.dataList(bucket)...)
	}
	oldBuckets := hm.OldBuckets()
	for _, bucket := range oldBuckets {
		if !evacuated(&bucket) {
			allData = append(allData, hm.dataList(bucket)...)
		}
	}
	return allData
}

func (hm *HashMap) ShowStat() {
	fmt.Printf("map \t flags=%d \t len=%d \t B=%d \t buckets=%d \t overflow=%d \t oldbuckets=%d \n", hm.Flags(), hm.Len(), hm.B(), hm.BucketCount(), hm.OverflowCount(), hm.OldBucketCount())
}

func (hm *HashMap) OldBucketCount() int {
	if hm.HMap().oldbuckets == nil {
		return 0
	}
	count := hm.BucketCount()
	if !(hm.Flags()&sameSizeGrow != 0) {
		count /= 2
	}
	return count
}

func (hm *HashMap) OldBucket(bucketIndex uintptr) bmap {
	return *(*bmap)(unsafe.Pointer(uintptr(hm.HMap().oldbuckets) + reflect.TypeOf(bmap{}).Size()*bucketIndex))
}

func (hm *HashMap) FirstOldBucket() bmap {
	return *(*bmap)(hm.HMap().oldbuckets)
}

func (hm *HashMap) OldBuckets() []bmap {
	count := hm.OldBucketCount()
	var buckets []bmap
	for i := 0; i < count; i++ {
		buckets = append(buckets, hm.OldBucket(uintptr(i)))
	}
	return buckets
}

func evacuated(b *bmap) bool {
	h := b.tophash[0]
	return h > emptyOne && h < minTopHash
}

func (hm *HashMap) sameSizeGrow() bool {
	return hm.HMap().flags&sameSizeGrow != 0
}
