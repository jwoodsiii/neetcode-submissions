func hasDuplicate(nums []int) bool {
    var cnts = make(map[int]int, len(nums))
	for _, i := range nums {
		cnts[i] += 1
		
		if cnts[i] > 1 {
			return true
		}
		fmt.Printf("%d: %d\n", i, cnts[i])
	}
	return false
}
