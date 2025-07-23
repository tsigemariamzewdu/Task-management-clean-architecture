package infrastruture



import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTServiceImpl implements domain.JWTService
type JWTServiceImpl struct {
	secretKey []byte
}

// NewJWTService returns a new instance of JWTServiceImpl with a secret key
func NewJWTService(secret string) *JWTServiceImpl {
	return &JWTServiceImpl{
		secretKey: []byte(secret),
	}
}


// GenerateToken creates a signed JWT token for the given user ID and role
func (j *JWTServiceImpl) GenerateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}


