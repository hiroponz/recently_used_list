package go_ginkgo

type RecentlyUsedList struct{}

var list [10]string
var count = 0

func (r *RecentlyUsedList) IsEmpty() bool {
	if count == 0 {
		return true
	} else {
		return false
	}
}

func (r *RecentlyUsedList) Push(str string) {
	list[count] = str
	count++
}

func (r *RecentlyUsedList) Fetch(index int) string {
	return list[index]
}
