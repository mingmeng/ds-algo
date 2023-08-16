package arraylist

func sortColors(nums []int) {
	zero := -1
	two := len(nums)
	for i := 0; i < two; {
		switch nums[i] {
		case 1:
			i++
		case 2:
			two -= 1
			nums[i], nums[two] = nums[two], nums[i]
		case 0:
			zero += 1
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		}
	}
}
