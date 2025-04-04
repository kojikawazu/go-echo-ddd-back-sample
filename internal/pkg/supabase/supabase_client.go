package pkg_supabase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pkg_logger "backend/internal/pkg/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Supabaseクライアント
type SupabaseClient struct {
	// Supabaseとのやり取りに使用するグローバルなコンテキスト。
	Ctx context.Context
	// Supabaseとの接続プールです。クエリ実行時に使用。
	Pool *pgxpool.Pool
}

// Supabaseクライアントのインスタンス化
func NewSupabaseClient() *SupabaseClient {
	return &SupabaseClient{
		Ctx: context.Background(),
	}
}

// Supabaseの接続を初期化
// Supabaseの接続URLを環境変数から取得し、コネクションプールを設定する。
// コネクションの最大数やアイドルタイム、シンプルプロトコルの使用を設定する。
// 成功時にはnilを返し、接続に失敗した場合はエラーメッセージを返す。
func (c *SupabaseClient) InitSupabase(logger *pkg_logger.AppLogger) error {
	logger.InfoLog.Println("Initializing Supabase client...")
	supabaseURL := os.Getenv("SUPABASE_URL") + "?sslmode=require"

	config, err := pgxpool.ParseConfig(supabaseURL)
	if err != nil {
		log.Printf("Unable to parse database URL: %v", err)
		return fmt.Errorf("unable to parse database URL: %v", err)
	}

	// コネクションプールの設定
	config.MaxConns = 10
	config.MaxConnIdleTime = 30 * time.Second
	// Prepared Statementの競合を防ぐためにSimple Protocolを優先
	config.ConnConfig.PreferSimpleProtocol = true

	logger.InfoLog.Println("Connecting supabase database...")
	c.Pool, err = pgxpool.ConnectConfig(c.Ctx, config)
	if err != nil {
		logger.ErrorLog.Printf("Unable to connect to Supabase: %v", err)
		return fmt.Errorf("unable to connect to Supabase: %v", err)
	}

	// 接続の確認
	logger.InfoLog.Println("Pinging supabase database...")
	err = c.Pool.Ping(c.Ctx)
	if err != nil {
		logger.ErrorLog.Printf("Unable to ping Supabase: %v", err)
		return fmt.Errorf("unable to ping Supabase: %v", err)
	}

	logger.InfoLog.Println("Connected to Supabase successfully")
	return nil
}

// Supabaseのコネクションプールをクローズ。
// この関数はアプリケーションのシャットダウン時に呼び出されることを想定する。
func (c *SupabaseClient) ClosePool(logger *pkg_logger.AppLogger) {
	if c.Pool != nil {
		c.Pool.Close()
		logger.InfoLog.Println("Supabase connection pool closed")
	}
}

// Supabaseに対してシンプルなクエリを実行し、接続が正しく動作しているかを確認する。
// クエリ結果として "1" を取得し、それをログに出力する。
// クエリに失敗した場合、エラーを返する。
func (c *SupabaseClient) TestQuery(logger *pkg_logger.AppLogger) error {
	logger.InfoLog.Println("Testing query...")
	query := `SELECT 1`
	rows, err := c.Pool.Query(c.Ctx, query)
	if err != nil {
		logger.ErrorLog.Printf("Failed to test query: %v", err)
		return err
	}
	logger.InfoLog.Println("Test query successful")
	defer rows.Close()

	for rows.Next() {
		var num int
		err := rows.Scan(&num)
		if err != nil {
			logger.ErrorLog.Printf("Failed to scan test query result: %v", err)
			return err
		}
		logger.InfoLog.Println("Test Query Result:", num)
	}

	logger.InfoLog.Println("Test query completed")
	return rows.Err()
}
