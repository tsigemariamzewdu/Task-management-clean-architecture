package infrastruture



import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordServiceImpl implements the domain.PasswordService interface
type PasswordServiceImpl struct{}

// hashes the given password using bcrypt.
func (p *PasswordServiceImpl) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

//  compares a hashed password with a plain text input.

func (p *PasswordServiceImpl) ComparePassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
