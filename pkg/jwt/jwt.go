package jwt

import (
	"fmt"
	"time"

	"gapp1/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

// Claims 定义 JWT 负载，包含用户 ID 和其他信息
type Claims struct {
	ID string `json:"id"` // 用户 ID
	jwt.RegisteredClaims
}

// GenerateJWT 生成 JWT token
func GenerateJWT(userID string) (string, error) {
	// 获取 JWT 配置信息
	secretKey := []byte(config.GetJWTSecretKey())
	issuer := config.GetJWTIssuer()
	expiration := time.Duration(config.GetJWTExpiration()) * time.Second

	// 创建 JWT 负载，使用用户 ID
	claims := Claims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		},
	}

	// 创建新的 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名并生成 token
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing the token: %v", err)
	}

	return signedToken, nil
}

// ParseJWT 解析 JWT token
func ParseJWT(tokenString string) (*Claims, error) {
	// 获取 JWT 密钥
	secretKey := []byte(config.GetJWTSecretKey())

	// 解析 token 并验证其签名
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保 token 使用 HMAC 签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// 如果解析成功且 token 有效
	if err != nil {
		return nil, fmt.Errorf("error parsing the token: %v", err)
	}

	// 提取并验证 JWT 负载
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

// VerifyToken 验证 JWT 是否有效
func VerifyToken(tokenString string) (bool, error) {
	_, err := ParseJWT(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ExtractToken 从请求头中提取 token
func ExtractToken(authorizationHeader string) string {
	// Authorization: Bearer <token>
	if len(authorizationHeader) > 7 && authorizationHeader[:7] == "Bearer " {
		return authorizationHeader[7:]
	}
	return ""
}
