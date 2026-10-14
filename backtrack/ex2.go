/*
 * 题目描述：
 * 给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。你可以按任意顺序返回答案。
 *
 * 示例：
 * - 输入: nums = [1, 2, 3]
 * - 输出: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
 * - 解释: 注意 [1,2,3] 和 [1,3,2] 被认为是不同的排列。
 */

package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	var length = len(nums)

	visited := make([]bool, length)

	var backtrack func()

	backtrack = func() {
		// return if archiving length(depth)
		if len(path) == length {
			tmp := make([]int, length)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// for each nums
		for i := 0; i < length; i++ {
			// skip elmt of used
			if visited[i] {
				continue
			}

			// append it
			visited[i] = true
			path = append(path, nums[i])

			// into next
			backtrack()

			// back it
			path = path[:len(path)-1]
			visited[i] = false
		}
	}

	backtrack()

	return res
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Printf("输出: %v\n", result)
}
