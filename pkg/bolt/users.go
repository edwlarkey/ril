package bolt

import (
	"time"

	"github.com/asdine/storm"
	"github.com/edwlarkey/ril/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *DB) InsertUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	user := models.User{
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
		Created:        time.Now(),
	}

	err = m.DB.Save(&user)

	if err != nil {
		return err
	}

	return nil
}

func (m *DB) AuthenticateUser(email, password string) (int, error) {
	user := models.User{}

	err := m.DB.One("Email", email, &user)
	if err == storm.ErrNotFound {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (m *DB) GetUser(id int) (*models.User, error) {
	user := models.User{}

	err := m.DB.One("ID", id, &user)
	if err == storm.ErrNotFound {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
