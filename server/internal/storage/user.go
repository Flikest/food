package storage

type User struct {
	ID          int    `json:"id"`
	Use         string `json:"use"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

func (s Storage) CreateUser() {
	createUserQuery := "INSERT INTO users(id, use, name, avatar, description) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.DB.Exec(s.Context, createUserQuery)
}

func (s Storage) UpdateUser() {
	updateUserQuery := "UPDATE users SET use=$2, name=$3, avatar=$4, description=$5 WHERE id=$1"
}

func (s Storage) DeleteUser() {
	deleteUserQuery := "DELETE FROM users WHERE id=$1"
}
