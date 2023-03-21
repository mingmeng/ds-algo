package arraylist

func removeElement(nums []int, val int) int {
	var fast, slow int

	for fast < len(nums)-1 {
		if nums[fast] == val {
			fast += 1
			continue
		}

		nums[slow] = nums[fast]
		slow++
		fast++
	}
	return slow
}
