package repositories

import (
	"swift/pkg/models"
)

func (r *Repository) GetUserFromDB(id int) (*models.User, error) {
	var u *models.User
	err := r.Db.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Repository) GetProfileFromDB(id int) (*models.User, error) {
	var profile *models.User
	err := r.Db.QueryRow(`SELECT * FROM profile WHERE user_id = $1`, id).Scan(&profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}




func (r *Repository) DeleteMyAccount(userID int) (err error) {
	
	_, err = r.Db.Exec(`update users set active = false, updated_at = current_timestamp
	 where id = $1`,userID)
	if err != nil {
		return err
	}
	return nil
}

func(r *Repository) UpdateProfileFromDB(id int, user *models.User)(*models.User,error){
	query :=`UPDATE users SET nik_name = $1,name = $2,birthday = $3, phone = $4,email = $5,password = $6 WHERE id = ?;` 
	_, err := r.Db.Exec(query, user.NikName, user.Name, user.Birthday, user.Phone, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil 
}
