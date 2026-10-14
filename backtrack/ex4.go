/*
 * 题目描述：
 *  给定一个 m x n 二维字符网格 board 和一个字符串单词 word。如果 word 存在于网格中，返回 true；否则，返回 false。
 *  单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中，相邻单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被多次使用。
 *
 * 示例：
 * - 输入:
 *    board = [
 *        ['A','B','C','E'],
 *        ['S','F','C','S'],
 *        ['A','D','E','E']
 *    ]
 *    word = "ABCCED"
 * - 输出：true
 * 解释：路径为 A(0,0) -> B(0,1) -> C(0,2) -> C(1,2) -> E(2,2) -> D(2,1)
 */

package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	rows, cols := len(board), len(board[0])

	// 定义移动方向：上，下，左，右
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	var dfs func(x, y, index int) bool
	dfs = func(x, y, index int) bool {
		// 剪枝
		if board[x][y] != word[index] {
			return false
		}

		// 全匹配
		if index == len(word)-1 {
			return true
		}

		// 提取字符
		temp := board[x][y]
		board[x][y] = '#'

		// 向4个方向探索
		for i := 0; i < 4; i++ {
			newX, newY := x+dx[i], y+dy[i]

			// 边界检查：确保新坐标在范围内
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
				if board[newX][newY] != '#' {
					if dfs(newX, newY, index+1) {
						return true
					}
				}
			}
		}

		// 撤销
		board[x][y] = temp
		return false
	}

	// 遍历网格中的每一个点，作为起点尝试探索
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Printf("输出: %v\n", exist(board, word))
	fmt.Printf("board路径：%c\n", board)
}
