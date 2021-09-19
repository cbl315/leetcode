package twoSum

func twoSum(nums []int, target int) []int {
	var minusMap = map[int]int{}
	for i, num := range nums {
		if j, ok := minusMap[num]; ok {
			return []int{j, i}
		}
		minus := target - num
		minusMap[minus] = i
	}
	return nil
}
