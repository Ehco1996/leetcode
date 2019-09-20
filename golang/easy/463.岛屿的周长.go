package easy

/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 *
 * https://leetcode-cn.com/problems/island-perimeter/description/
 *
 * algorithms
 * Easy (61.98%)
 * Likes:    117
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 12.9K
 * Testcase Example:  '[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]'
 *
 * 给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
 *
 * 网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
 *
 * 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100
 * 。计算这个岛屿的周长。
 *
 *
 *
 * 示例 :
 *
 * 输入:
 * [[0,1,0,0],
 * ⁠[1,1,1,0],
 * ⁠[0,1,0,0],
 * ⁠[1,1,0,0]]
 *
 * 输出: 16
 *
 * 解释: 它的周长是下面图片中的 16 个黄色的边：
 *
 *
 *
 *
 */

type Point struct {
	x       int
	y       int
	number  int
	visited bool
	isLand  bool
}

func NewPoint(x, y int, isLand bool) *Point {
	return &Point{x: x, y: y, isLand: isLand}
}

type Boards struct {
	Hash map[string]*Point
	Grid [][]int
}

func (b *Boards) GetOrCreatePoint(x, y int) *Point {
	key := string(x) + "-" + string(y)

	p, ok := b.Hash[key]
	if ok {
		return p
	}
	if 0 <= x && x < len(b.Grid) && 0 <= y && y < len(b.Grid[x]) {
		b.Hash[key] = NewPoint(x, y, b.Grid[x][y] == 1)
	} else {
		b.Hash[key] = NewPoint(x, y, false)
	}
	return b.Hash[key]
}

func (b *Boards) Visit(p *Point) {
	if p.visited {
		return
	} else {
		p.visited = true
		if !p.isLand {
			return
		}
		up := b.GetOrCreatePoint(p.x, p.y-1)
		if !up.isLand {
			p.number++
		}

		right := b.GetOrCreatePoint(p.x+1, p.y)
		if !right.isLand {
			p.number++
		}

		down := b.GetOrCreatePoint(p.x, p.y+1)
		if !down.isLand {
			p.number++
		}

		left := b.GetOrCreatePoint(p.x-1, p.y)
		if !left.isLand {
			p.number++
		}
	}

}

func islandPerimeter(grid [][]int) int {
	// 暴力算的
	// 但可以这样解 由于岛屿内没有湖,所以只需要求出 北面(或南面) + 西面(或东面)的长度再乘2即可
	b := &Boards{
		Hash: make(map[string]*Point),
		Grid: grid,
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			p := b.GetOrCreatePoint(x, y)
			if p != nil {
				b.Visit(p)
			}
		}
	}

	res := 0
	for _, v := range b.Hash {
		res += v.number
	}
	return res
}
