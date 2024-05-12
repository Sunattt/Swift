package repositories

import (
	"database/sql"
	"log"
	"swift/pkg/db"
	"swift/pkg/models"
	"time"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Db: db,
	}
}

func (r *Repository) IsNikNameFree(nikName string) (isFree bool, err error) {
	var nameDB int

	err = r.Db.QueryRow(`
	SELECT COUNT(*)
	FROM users WHERE nik_name = $1`, nikName).Scan(&nameDB)
	if err != nil {
		return false, err
	}

	if nameDB != 0 {
		return true, nil
	}
	return false, nil
}

func (r *Repository) AddUserToDB(u *models.User, hash string) error {
	now := time.Now()
	u.CreatedAt = now
	query := `INSERT INTO users(nik_name, name, phone, birthday, email, password, create_at)
	VALUES (?, ?, ?, ?, ?, ?, ?);`
	_, err := r.Db.Exec(query, u.NikName, u.Name, u.Phone, u.Birthday, u.Email, hash, u.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func(r *Repository ) AddTokenToDb(userId int ,token string)error{
	var affacted int64 
	err := r.Db.QueryRow(`SELECT COUNT(*) FROM tokens WHERE user_id = $1`, userId).Scan(&affacted)
	if err != nil {
		return err
	}

	if affacted == 0{
		_, err = r.Db.Exec(`INSERT INTO tokens(token, user_id) VALUES($1, $2)`, token, userId)
		if err != nil {
return err 
		}
	}else{
		_, err = r.Db.Exec(`UPDATE tokens SET token = $1, update_at = $2 WHERE user_id = $3`, token, time.Now(), userId)
		if err != nil {
			return err
		}
	}
	return nil
}


func (r *Repository) DBCheckNikName(nik string) (user *models.User, userCount int8, err error) {
	err = db.DB.QueryRow(`
	SELECT COUNT(*)
	FROM users WHERE nik = $1 and active = true`, nik).Scan(&userCount)
	if err != nil {
		log.Println("[ERROR] during check nik and pass !!!")
		return
	}
	return
}

func (r *Repository) DBCheckActiveById(id int) (active bool, err error) {
	err = db.DB.QueryRow(`
	SELECT active FROM users
	WHERE id = $1`, id).Scan(&active)
	if err != nil {
		return false, err
	}
	return active, nil
}
