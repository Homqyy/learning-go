/*
 * 题目描述：
 * 给你一个 无重复元素 的正整数数组 candidates 和一个目标整数 target，找出 candidates 中可以使数字和为目标数 target 的所有不同组合，并以二维切片的形式返回。你可以按 任意顺序 返回这些组合。
 * candidates 中的 同一个数字可以无限制重复被选取。
 *
 * 示例：
 * - 输入: candidates = [2, 3, 6, 7], target = 7
 * - 输出: [[2, 2, 3], [7]]
 * - 解释: 2 和 3 可以形成一组候选，2 + 2 + 3 = 7。注意 2 可以使用多次。7 也是一个候选，7 = 7。
 */

package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(start int, currentTarget int)

	/*
	 * @start 起始索引
	 * @currentTarget 当前目标值
	 */
	backtrack = func(start int, currentTarget int) {
		if currentTarget == 0 {
			// 找到一个符合条件的组合，加入结果集
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		if currentTarget < 0 {
			// 当前组合不符合条件，直接返回
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, currentTarget-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return res
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Printf("输出: %v\n", result)
}
