package services

import (
	"database/sql"

	"github.com/Aniketyadav44/go-backend-template/internal/models"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	query := "select id, username, email, created_at from users"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
