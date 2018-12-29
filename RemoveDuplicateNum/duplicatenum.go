package duplicatenum

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

// 示例 1:
// 给定数组 nums = [1,1,2],
// 函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
// 你不需要考虑数组中超出新长度后面的元素。

// 示例 2:
// 给定 nums = [0,0,1,1,1,2,2,3,3,4],
// 函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
// 你不需要考虑数组中超出新长度后面的元素。

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	newLength := 0
	for i := 1; i < length; i++ {
		if nums[newLength] != nums[i] {
			newLength++
			nums[newLength] = nums[i]
		}
	}
	return newLength + 1
}

// 适用于无排序数组
// func removeDuplicates(nums []int) int {
// 	for i := 1; i < len(nums); i++ {
// 		num := nums[i]
// 		array := nums[:i]
// 		for _, temp := range array {
// 			if num == temp {
// 				if i >= len(nums) {
// 					nums = nums[:i]
// 					i--
// 				} else {
// 					nums = append(nums[:i], nums[i+1:]...)
// 					i--
// 				}
// 			}
// 		}
// 	}
// 	return len(nums)
// }
