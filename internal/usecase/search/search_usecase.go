package search_usecase

import (
	pkg_logger "backend/internal/pkg/logger"
)

// Searchユースケース(IF)
type ISearchUsecase interface {
	LinearSearch(arr []int, target int) int
	BinarySearch(arr []int, target int) int
	BFS(graph [][]int) int
	DFS(graph [][]int) bool
}

// Searchユースケース(Impl)
type SearchUsecase struct {
	Logger *pkg_logger.AppLogger
}

// Searchユースケースのインスタンス化
func NewSearchUsecase(l *pkg_logger.AppLogger) ISearchUsecase {
	return &SearchUsecase{
		Logger: l,
	}
}

// 線形探索
// 配列の中にtargetが存在するかを探す
// [採用パターン]
// - 小規模データ（n≦10⁴）
// - 未ソート
// - 部分一致検索など
func (u *SearchUsecase) LinearSearch(arr []int, target int) int {
	index := -1

	for i, num := range arr {
		if num == target {
			index = i
			break
		}
	}

	return index
}

// 二分探索
// 配列の中にtargetが存在するかを探す
// [採用パターン]
// - ソート済み
// - 単調増加/減少の条件での最大/最小探索
func (u *SearchUsecase) BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// BFS（幅優先探索）
// 迷路探索
// [採用パターン]
// - 最短経路（非加重）
// - 迷路/レベル探索
// - 伝播問題
func (u *SearchUsecase) BFS(graph [][]int) int {
	// 座標と距離
	type Point struct {
		X, Y, Dist int
	}

	// 幅、高さ
	H, W := len(graph), len(graph[0])
	// 上下左右
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	// キュー
	queue := []Point{{0, 0, 1}}

	// 初期化
	visited := make([][]bool, H)
	for i := range visited {
		visited[i] = make([]bool, W)
	}
	// 訪問済み
	visited[0][0] = true

	// キューが空になるまでループ
	for len(queue) > 0 {

		p := queue[0]
		queue = queue[1:]

		if p.X == H-1 && p.Y == W-1 {
			// ゴールに到達したら距離を返す
			return p.Dist
		}

		// 上下左右を探索
		for _, d := range dirs {
			nx, ny := p.X+d[0], p.Y+d[1]
			if nx >= 0 && ny >= 0 && nx < H && ny < W && !visited[nx][ny] && graph[nx][ny] == 0 {
				// 訪問済みにする
				visited[nx][ny] = true
				// キューに追加
				queue = append(queue, Point{nx, ny, p.Dist + 1})
			}
		}
	}

	// ゴールに到達できなかったら-1を返す
	return -1
}

// DFS（深さ優先探索）
// 迷路探索
// [採用パターン]
// - 全探索
// - 組み合わせ列挙
// - 到達可能性の確認
func (u *SearchUsecase) DFS(graph [][]int) bool {
	H := len(graph)
	W := len(graph[0])

	visited := make([][]bool, H)
	for i := range visited {
		visited[i] = make([]bool, W)
	}

	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		// 境界外チェック
		if x < 0 || y < 0 || x >= H || y >= W {
			return false
		}
		// 壁チェック、訪問済みチェック
		if graph[x][y] == 1 || visited[x][y] {
			return false
		}
		// ゴール到達チェック
		if x == H-1 && y == W-1 {
			return true
		}

		// 訪問済みにする
		visited[x][y] = true

		// 上下左右を再帰的に探索
		return dfs(x+1, y) || dfs(x-1, y) || dfs(x, y+1) || dfs(x, y-1)
	}

	// スタート地点から探索
	return dfs(0, 0)
}
