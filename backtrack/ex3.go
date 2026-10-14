/*
 * 题目描述：
 * 给定一个可包含重复数字的序列 nums，按任意顺序返回所有不重复的全排列。
 *
 * 示例：
 * - 输入: nums = [1, 2, 1]
 * - 输出: [[1,1,2], [1,2,1], [2,1,1]]
 * - 解释: 解释: 注意，虽然有两个 1，但 [1,1,2] 只能出现一次，不能因为有两个 1 就生成两个相同的 [1,1,2]。
 */

package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	var length = len(nums)

	sort.Ints(nums)

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

			// append it if it's not a duplicate in the current depth
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

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
	nums := []int{1, 2, 1}
	result := permute(nums)
	fmt.Printf("输出: %v\n", result)
}
