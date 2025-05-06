package search_handler

import (
	pkg_logger "backend/internal/pkg/logger"
	usecase_search "backend/internal/usecase/search"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Todoハンドラ(Impl)
type SearchHandler struct {
	Logger        *pkg_logger.AppLogger
	searchUsecase usecase_search.ISearchUsecase
}

// Todoハンドラのインスタンス化
func NewSearchHandler(l *pkg_logger.AppLogger, su usecase_search.ISearchUsecase) *SearchHandler {
	return &SearchHandler{
		Logger:        l,
		searchUsecase: su,
	}
}

// 線形探索
func (h *SearchHandler) LinearSearch(c echo.Context) error {
	h.Logger.InfoLog.Println("LinearSearch called")

	// リクエストボディ
	body := struct {
		Arr    []int `json:"arr"`
		Target int   `json:"target"`
	}{}
	if err := c.Bind(&body); err != nil {
		h.Logger.ErrorLog.Println("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	h.Logger.InfoLog.Println("request arr: ", body.Arr)
	h.Logger.InfoLog.Println("request target: ", body.Target)

	// 線形探索を実行
	index := h.searchUsecase.LinearSearch(body.Arr, body.Target)

	// 結果をJSON形式で返す
	h.Logger.InfoLog.Println("index: ", index)
	return c.JSON(http.StatusOK, map[string]int{"index": index})
}

// 二分探索
func (h *SearchHandler) BinarySearch(c echo.Context) error {
	h.Logger.InfoLog.Println("BinarySearch called")

	// リクエストボディ
	body := struct {
		Arr    []int `json:"arr"`
		Target int   `json:"target"`
	}{}
	if err := c.Bind(&body); err != nil {
		h.Logger.ErrorLog.Println("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	h.Logger.InfoLog.Println("request arr: ", body.Arr)
	h.Logger.InfoLog.Println("request target: ", body.Target)

	// 二分探索を実行
	index := h.searchUsecase.BinarySearch(body.Arr, body.Target)

	// 結果をJSON形式で返す
	h.Logger.InfoLog.Println("index: ", index)
	return c.JSON(http.StatusOK, map[string]int{"index": index})
}

// BFS（幅優先探索）
func (h *SearchHandler) BFS(c echo.Context) error {
	h.Logger.InfoLog.Println("BFS called")

	// リクエストボディ
	body := struct {
		Graph [][]int `json:"graph"`
	}{}
	if err := c.Bind(&body); err != nil {
		h.Logger.ErrorLog.Println("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	h.Logger.InfoLog.Println("request graph: ", body.Graph)

	// BFSを実行
	steps := h.searchUsecase.BFS(body.Graph)

	// 結果をJSON形式で返す
	h.Logger.InfoLog.Println("steps: ", steps)
	return c.JSON(http.StatusOK, map[string]int{"steps": steps})
}

// DFS（深さ優先探索）
func (h *SearchHandler) DFS(c echo.Context) error {
	h.Logger.InfoLog.Println("DFS called")

	// リクエストボディ
	body := struct {
		Graph [][]int `json:"graph"`
	}{}
	if err := c.Bind(&body); err != nil {
		h.Logger.ErrorLog.Println("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	h.Logger.InfoLog.Println("request graph: ", body.Graph)

	// DFSを実行
	result := h.searchUsecase.DFS(body.Graph)

	// 結果をJSON形式で返す
	h.Logger.InfoLog.Println("result: ", result)
	return c.JSON(http.StatusOK, map[string]bool{"result": result})
}
