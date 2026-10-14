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

import (
	"fmt"
)

func resolveQueen(n int) [][]string {
	var res [][]string

	queues := make([]int, n) // 记录行的皇后位置
	for i := range queues {
		// 初始化为 -1
		queues[i] = -1
	}

	// 记录占用的位置
	cols := make([]bool, n)      // 记录哪列已经被占用
	diag1 := make([]bool, 2*n-1) // 记录主对角被占用的位置
	diag2 := make([]bool, 2*n-1) // 记录副对角被占用的位置

	// TODO: 定义回溯
	var backtrack func(row int)
	backtrack = func(row int) {
		// 结束条件：当达n时，说明已经有符合条件的位置组了
		if row == n {
			// 生成占位结果
			board := make([]string, n)
			for i := range queues {
				pos := queues[i]

				// 构建一个字符串，默认全部为 '.'
				s := make([]byte, n)
				for j := range s {
					s[j] = '.'
				}
				s[pos] = 'Q' // 记录皇后位置

				// 写入结果
				board[i] = string(s)
			}

			res = append(res, board)
			return
		}

		// 遍历所有列
		for col := 0; col < n; col++ {
			// 裁枝：如果位置已经被占用，则跳过
			if cols[col] || diag1[row-col+n-1] || diag2[row+col] {
				continue
			}

			// 选择
			queues[row] = col
			cols[col] = true
			diag1[row-col+n-1] = true
			diag2[row+col] = true

			// 进入下一行
			backtrack(row + 1)

			// 回溯
			queues[row] = -1
			cols[col] = false
			diag1[row-col+n-1] = false
			diag2[row+col] = false
		}
	}

	// 从第0行开始
	backtrack(0)

	// 返回结果
	return res
}

func main() {
	n := 4
	result := resolveQueen(n)

	// 打印结果
	for _, board := range result {
		for _, row := range board {
			fmt.Printf("%s\n", row)
		}
		fmt.Println("---")
	}
}
