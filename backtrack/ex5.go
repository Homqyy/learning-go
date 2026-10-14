/*
 * 题目描述：
 *   按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
 *   n 皇后问题研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *   给你一个整数 n，返回所有不同的 n 皇后问题 的解决方案。每一种解法包含一个不同的棋子放置方案，其中字符串 'Q' 和 '.' 分别代表皇后和空位。
 *
 * 示例：
 * - 输入: n = 4
 * - 输出:
 *    [
 *      [".Q..", "...Q", "Q...", "..Q."],
 *      ["..Q.", "Q...", "...Q", ".Q.."]
 *    ]
 */

package main

import "fmt"

func solveQueens(n int) [][]string {
	var res [][]string

	// queens 数组记录每一行皇后放在了哪一列
	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}

	// 记录冲突状态的数组
	cols := make([]bool, n)      // 记录列是否被占用
	diag1 := make([]bool, 2*n-1) // 记录主对角线是否被占用
	diag2 := make([]bool, 2*n-1) // 记录副对角线是否被占用

	// 定义回溯函数，按行进行放置
	var backtrack func(row int)
	backtrack = func(row int) {
		// TODO: 结束条件：如果成功放置到了第n行，说明找到了一组合法解
		if row == n {
			// 构造字符串
			board := make([]string, n)
			for i := 0; i < n; i++ {
				b := make([]byte, n)
				for j := 0; j < n; j++ {
					b[j] = '.'
				}
				b[queens[i]] = 'Q'
				board[i] = string(b)
			}
			res = append(res, board)
			return
		}

		// 遍历当前行的所有列，尝试放置皇后
		for col := 0; col < n; col++ {
			// 剪枝：检查是否有冲突
			if cols[col] || diag1[row-col+n-1] || diag2[row+col] {
				continue
			}

			// 选择
			queens[row] = col
			cols[col] = true
			diag1[row-col+n-1] = true
			diag2[row+col] = true

			// 进入下一行
			backtrack(row + 1)

			// 回溯
			cols[col] = false
			diag1[row-col+n-1] = false
			diag2[row+col] = false
		}
	}

	// 从第0行开始
	backtrack(0)
	return res
}

func main() {
	n := 4
	result := solveQueens(n)

	// 打印有效的path
	for _, board := range result {
		for _, row := range board {
			fmt.Println(row)
		}
		fmt.Println("---")
	}
}
