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

import (
	"fmt"
	"slices"
)

func wordMatch(board [][]rune, wordString string) bool {
	var path []rune
	var matched bool
	xMax := len(board)
	yMax := len(board[0])

	word := []rune(wordString)

	var backtrack func(int, int)

	backtrack = func(x int, y int) {
		fmt.Printf("x=%v, y=%v\n", x, y)
		// not match next word
		pos := len(path)
		if pos > 0 && path[pos-1] != word[pos-1] {
			fmt.Println("no match next word")
			return
		}

		// overflow
		if x >= xMax || y >= yMax {
			fmt.Printf("x or y overflow: %v, %v\n", x, y)
			return
		}

		// full match word
		if slices.Equal(word, path) {
			fmt.Printf("full match word: %v\n", string(path))
			matched = true
			return
		}

		// pick it: first to pick x, next to pick y
		fmt.Printf("path is %s, try word(x): %c\n", string(path), board[x][y])
		path = append(path, board[x][y])
		backtrack(x+1, y)

		if matched {
			return
		}

		path = path[:len(path)-1]

		backtrack(x, y+1)

		if matched {
			return
		}

		path = path[:len(path)-1]
	}

	backtrack(0, 0)

	return matched
}

func main() {
	board := [][]rune{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCCED"

	result := wordMatch(board, word)

	fmt.Printf("%v\n", result)
}
