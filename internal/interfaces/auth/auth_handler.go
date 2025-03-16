package interfaces_auth

import (
	pkg_logger "backend/internal/pkg/logger"
	usecase_auth "backend/internal/usecase/auth"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// 認証ハンドラ(Impl)
type AuthHandler struct {
	Logger      *pkg_logger.AppLogger
	authUsecase usecase_auth.IAuthUsecase
}

// 認証ハンドラのインスタンス化
func NewAuthHandler(l *pkg_logger.AppLogger, u usecase_auth.IAuthUsecase) *AuthHandler {
	return &AuthHandler{
		Logger:      l,
		authUsecase: u,
	}
}

// ログイン
func (h *AuthHandler) Login(c echo.Context) error {
	h.Logger.InfoLog.Println("Login called")

	// ログインリクエストボディを取得
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// リクエストボディをパース
	if err := c.Bind(&loginRequest); err != nil {
		h.Logger.ErrorLog.Printf("Failed to parse login request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// ログイン(usecase層)
	id, err := h.authUsecase.Login(loginRequest.Email, loginRequest.Password)
	// エラーハンドリング
	if err != nil {
		switch err {
		case errors.New("invalid email or password"):
			h.Logger.ErrorLog.Println("Invalid email or password")
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
		case errors.New("invalid email format"):
			h.Logger.ErrorLog.Println("Invalid email format")
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid email format"})
		case errors.New("failed to login"):
			h.Logger.ErrorLog.Println("Failed to login")
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to login"})
		default:
			h.Logger.ErrorLog.Printf("Failed to get all users: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	}

	// JWTトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": "user",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// JWTトークンをシグネーション
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		h.Logger.ErrorLog.Printf("Failed to sign token: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to sign token"})
	}

	h.Logger.InfoLog.Println("Login successful. 1 user found")
	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

// 認証ミドルウェア
func (h *AuthHandler) AuthorizationMiddleware(next echo.HandlerFunc, requiredRole string) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing Authorization header"})
		}

		// "Bearer " を取り除く
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// JWT をパース
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signing method")
			}
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}

		// クレームからユーザーIDとロールを取得
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token claims"})
		}

		// ロールを確認（例: "admin", "user" など）
		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "Insufficient permissions"})
		}

		// ユーザーIDをコンテキストに保存
		c.Set("userId", claims["id"])
		c.Set("role", role)

		return next(c)
	}
}
