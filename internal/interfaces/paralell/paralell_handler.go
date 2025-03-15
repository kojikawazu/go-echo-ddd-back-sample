package interfaces_paralell

import (
	"backend/config"
	pkg_logger "backend/internal/pkg/logger"
	"backend/utils"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

// 並列動作のハンドラ
type ParalellHandler struct {
	AppConfig *config.AppConfig
	Logger    *pkg_logger.AppLogger
}

// 並列動作のインスタンス化
func NewParalellHandler(appConfig *config.AppConfig, logger *pkg_logger.AppLogger) *ParalellHandler {
	return &ParalellHandler{
		AppConfig: appConfig,
		Logger:    logger,
	}
}

// 並列動作のサンプル
func (h *ParalellHandler) ExecParallel(c echo.Context) error {
	h.Logger.InfoLog.Println("ParallelHandler started")

	// 開始時間を計測
	start := time.Now()

	urls := []string{
		h.AppConfig.TestAPI + "/posts",
		h.AppConfig.TestAPI + "/comments",
		h.AppConfig.TestAPI + "/albums",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(url string, i int) {
			defer wg.Done()
			h.Logger.InfoLog.Println("Request", i, "started")
			data := utils.FetchAPI(url)
			results <- data
		}(url, i)
	}

	wg.Wait()
	close(results)

	var allResults []string
	for data := range results {
		allResults = append(allResults, data)
	}

	// 終了時間を計測
	end := time.Now()

	h.Logger.InfoLog.Println("ParallelHandler completed")
	h.Logger.InfoLog.Println("Time taken:", end.Sub(start))

	return c.String(http.StatusOK, strings.Join(allResults, "\n"))
}

// 逐次処理のサンプル
func (h *ParalellHandler) ExecSeries(c echo.Context) error {
	h.Logger.InfoLog.Println("SeriesHandler started")

	// 開始時間を計測
	start := time.Now()

	urls := []string{
		h.AppConfig.TestAPI + "/posts",
		h.AppConfig.TestAPI + "/comments",
		h.AppConfig.TestAPI + "/albums",
	}
	var allResults []string

	for i, url := range urls {
		h.Logger.InfoLog.Println("Request", i, "started")
		data := utils.FetchAPI(url)
		allResults = append(allResults, data)
	}

	// 終了時間を計測
	end := time.Now()

	h.Logger.InfoLog.Println("SeriesHandler completed")
	h.Logger.InfoLog.Println("Time taken:", end.Sub(start))

	return c.String(http.StatusOK, strings.Join(allResults, "\n"))
}
