package go_ginkgo

type RecentlyUsedList struct {
	list  [10]string
	count int
}

func (r *RecentlyUsedList) IsEmpty() bool {
	return r.count == 0
}

func (r *RecentlyUsedList) Push(str string) {
	if r.Index(str) == -1 {
		r.list[r.count] = str
		r.count++
	}
}

func (r *RecentlyUsedList) Fetch(index int) string {
	return r.list[r.count-1-index]
}

func (r *RecentlyUsedList) Count() int {
	return r.count
}

func (r *RecentlyUsedList) Index(str string) int {
	for i := 0; i < r.count; i++ {
		if r.Fetch(i) == str {
			return i
		}
	}
	return -1
}
