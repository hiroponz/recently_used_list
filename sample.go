package go_ginkgo

type RecentlyUsedList struct {
	list  [10]string
	count int
}

func (r *RecentlyUsedList) IsEmpty() bool {
	if r.count == 0 {
		return true
	} else {
		return false
	}
}

func (r *RecentlyUsedList) Push(str string) {
	r.list[r.count] = str
	r.count++
}

func (r *RecentlyUsedList) Fetch(index int) string {
	return r.list[r.count-1-index]
}

func (r *RecentlyUsedList) Count() int {
	return r.count
}
