package database

import (
	"fmt"
	"golang/calendar/entities"
	"golang/calendar/infrastructure/database"
	"log"
	"strconv"
)

type NikkiRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *NikkiRepository) FindAll() (entities.Nikkis, error) {
	var nikkis entities.Nikkis
	fmt.Println("show nikkis")
	rows, err := repo.SqlHandler.DB.Query("SELECT * from nikkis;")
	if err != nil {
		log.Print("error executing database query: ", err)
	}
	defer rows.Close()

	var nikkis_table_colum Nikkis_table
	for rows.Next() {
		var nikki entities.Nikki
		err := rows.Scan(
			&nikkis_table_colum.ID,
			&nikkis_table_colum.UserID,
			&nikkis_table_colum.Date,
			&nikkis_table_colum.Title,
			&nikkis_table_colum.Content,
			&nikkis_table_colum.NumberOfPhotos,
			&nikkis_table_colum.CreatedAt,
			&nikkis_table_colum.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		nikki.ID = nikkis_table_colum.ID
		nikki.UserID = nikkis_table_colum.UserID
		nikki.Title = nikkis_table_colum.Title
		nikki.Date = nikkis_table_colum.Date
		nikki.Content = nikkis_table_colum.Content
		nikki.NumberOfPhotos = nikkis_table_colum.NumberOfPhotos
		nikkis = append(nikkis, nikki)
	}
	return nikkis, nil
}
func (repo *NikkiRepository) FindNikki(UserId int, Date int) (entities.Nikki, error) {
	var nikkis_table_colum Nikkis_table
	var nikki entities.Nikki
	nikki.ID = 0

	rows, err := repo.SqlHandler.DB.Query("SELECT * FROM nikkis WHERE user_id = ? and date = ? LIMIT 1;", UserId, Date)
	if err != nil {
		log.Print("error executing database query: ", err)
		return nikki, err
	}
	for rows.Next() {
		err := rows.Scan(
			&nikkis_table_colum.ID,
			&nikkis_table_colum.UserID,
			&nikkis_table_colum.Date,
			&nikkis_table_colum.Title,
			&nikkis_table_colum.Content,
			&nikkis_table_colum.NumberOfPhotos,
			&nikkis_table_colum.CreatedAt,
			&nikkis_table_colum.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
			return nikki, nil
		}
		nikki.ID = nikkis_table_colum.ID
		nikki.UserID = nikkis_table_colum.UserID
		nikki.Title = nikkis_table_colum.Title
		nikki.Date = nikkis_table_colum.Date
		nikki.Content = nikkis_table_colum.Content
		nikki.NumberOfPhotos = nikkis_table_colum.NumberOfPhotos
		fmt.Println("find nikki userId = " + strconv.Itoa(nikkis_table_colum.UserID) + " date = " + strconv.Itoa(nikkis_table_colum.Date))
		// nikkis = append(nikkis, nikki)
	}
	fmt.Println(nikki)
	fmt.Println(nikkis_table_colum)
	return nikki, nil
}
func (repo *NikkiRepository) FindPhotos(NikkiID int, NumberOfPhotos int) (entities.Photos, error) {
	var Photos entities.Photos
	rows, err := repo.SqlHandler.DB.Query("SELECT id,photo FROM photos WHERE nikki_id = ? LIMIT ?;",
		NikkiID, NumberOfPhotos)
	if err != nil {
		log.Print("error executing database query: ", err)
		return Photos, err
	}
	for rows.Next() {
		var photo entities.Photo
		err := rows.Scan(&photo.ID, &photo.Photo)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
			return Photos, err
		}
		Photos = append(Photos, photo)
	}
	return Photos, nil
}
func (repo *NikkiRepository) CreateNikki(UserId int, Date int, Title string, Content string, NumberOfPhotos int) (entities.Nikki, error) {
	var nikki entities.Nikki
	statement := "INSERT INTO nikkis(user_id,date,title,content,number_of_photos) VALUES(?,?,?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare(statement) error")
		return nikki, err
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(UserId, Date, Title, Content, NumberOfPhotos)
	if err != nil {
		fmt.Println("stmtInsert.Exec　error")
		return nikki, err
	}
	lastInsertID, err := result.LastInsertId()

	err = repo.SqlHandler.DB.QueryRow("SELECT id,user_id,date,title,content,number_of_photos FROM nikkis WHERE id = ?", lastInsertID).Scan(&nikki.ID, &nikki.UserID,
		&nikki.Date, &nikki.Title, &nikki.Content, &nikki.NumberOfPhotos)
	if err != nil {
		log.Fatal(err)
	}

	return nikki, nil
}
func (repo *NikkiRepository) DeleteNikki(UserID int, Date int) (int, int, int, error) {
	stmtDelete, err := repo.SqlHandler.DB.Prepare("DELETE FROM nikkis WHERE user_id = ? and date = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(UserID, Date)
	if err != nil {
		panic(err.Error())
	}

	_rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(_rowsAffect)
	rowsAffect := int(_rowsAffect)
	if rowsAffect == 0 {
		fmt.Println("削除できません")
	} else if rowsAffect == 1 {
		fmt.Println("削除完了")
	} else {
		fmt.Println("DB table エラー") //削除データが2個以上は起らないはず
	}
	return UserID, Date, rowsAffect, err
}
func (repo *NikkiRepository) EditNikki(UserId int, Date int, Title string, Content string) {
	stmtEdit, err := repo.SqlHandler.DB.Prepare("UPDATE nikkis SET title = ?,content = ? WHERE user_id = ? and date = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtEdit.Close()

	result, err := stmtEdit.Exec(Title, Content, 2, 20191128)
	if err != nil {
		panic(err.Error())
	}
	_rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(_rowsAffect)
	rowsAffect := int(_rowsAffect)
	if rowsAffect == 0 {
		fmt.Println("削除できません")
	} else if rowsAffect == 1 {
		fmt.Println("削除完了")
	} else {
		fmt.Println("DB table エラー") //削除データが2個以上は起らないはず
	}
}
func (repo *NikkiRepository) InsertPhoto(NikkiId int, UserId int, Date int, PhotoId int, Photo string) {
	fmt.Println("RegisterPhoto")
	statement := "INSERT INTO photos(nikki_id,user_id,date,photo_id,photo) VALUES(?,?,?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare(statement) error")
	}
	defer stmtInsert.Close()

	res, err := stmtInsert.Exec(NikkiId, UserId, Date, PhotoId, Photo)
	if err != nil {
		fmt.Println(res)
		fmt.Println("stmtInsert.Exec　error")
	}
}

// func (repo *NikkiRepository) FindAllPhotos(){
// 	var photos entities.Photos
// 	fmt.Println("FindAllPhotos")
// 	rows, err := repo.SqlHandler.DB.Query("SELECT * from photos;")
// 	if err != nil {
// 		log.Print("error executing database query: ", err)
// 	}
// 	defer rows.Close()

// 	var photos_table_colum Photos_table
// 	for rows.Next() {
// 		var photo entities.Photo
// 		err := rows.Scan(
// 			&photos_table_colum.ID,
// 			&photos_table_colum.NikkiID,
// 			&photos_table_colum.UserID,
// 			&photos_table_colum.Date,
// 			&photos_table_colum.PhotoId,
// 			&photos_table_colum.Photo,
// 			&photos_table_colum.CreatedAt,
// 			&photos_table_colum.UpdatedAt)
// 		if err != nil {
// 			fmt.Println(err)
// 			panic(err.Error())
// 		}
// 		nikki.ID = nikkis_table_colum.ID
// 		nikki.UserID = nikkis_table_colum.UserID
// 		nikki.Title = nikkis_table_colum.Title
// 		nikki.Date = nikkis_table_colum.Date
// 		nikki.Content = nikkis_table_colum.Content
// 		nikkis = append(nikkis, nikki)
// 	}
// 	return nikkis, nil
// }
