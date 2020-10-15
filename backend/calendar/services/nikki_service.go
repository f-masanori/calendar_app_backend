package services

import (
	"fmt"
	"golang/calendar/entities"
)

type NikkiRepository interface {
	FindAll() (entities.Nikkis, error)
	FindNikki(int, int) (entities.Nikki, error)
	CreateNikki(int, int, string, string, int) (entities.Nikki, error)
	DeleteNikki(int, int) (int, int, int, error)
	EditNikki(int, int, string, string)
	InsertPhoto(int, int, int, int, string)
	FindPhotos(int, int) (entities.Photos, error)
}
type NikkiService struct {
	NikkiRepository NikkiRepository
}

/* nikki delete時に使用 */
type ConfirmDelete struct {
	UserId     int
	Date       int
	RowsAffect int
	Err        error
}

func (n *NikkiService) GetAll() (entities.Nikkis, error) {

	nikkis, err := n.NikkiRepository.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nikkis)

	return nikkis, err
}

func (n *NikkiService) CreateNikki(UserId int, Date int, Title string, Content string, NumberOfPhotos int) (entities.Nikki, error) {
	nikki, err := n.NikkiRepository.CreateNikki(UserId, Date, Title, Content, NumberOfPhotos)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nikki)
	return nikki, err
}

func (n *NikkiService) EditNikki(UserId int, Date int, Title string, Content string) {
	n.NikkiRepository.EditNikki(UserId, Date, Title, Content)
}

func (n *NikkiService) DeleteNikki(UserID int, Date int) *ConfirmDelete {

	ConfirmDelete := new(ConfirmDelete)
	ConfirmDelete.UserId, ConfirmDelete.Date, ConfirmDelete.RowsAffect, ConfirmDelete.Err = n.NikkiRepository.DeleteNikki(UserID, Date)
	if ConfirmDelete.Err != nil {
		fmt.Println(ConfirmDelete.Err)
	}
	return ConfirmDelete
}

func (n *NikkiService) GetNikki(UserID int, Date int) (entities.Nikki, error) {
	nikki, err := n.NikkiRepository.FindNikki(UserID, Date)
	if err != nil {
		fmt.Println(err)
		//ここでreturnして止める
	}
	if nikki.ID == 0 {
		fmt.Println("日記は存在しません")
		//ここでreturnして止める
	}
	if nikki.NumberOfPhotos != 0 {
		fmt.Println("写真が存在")
		photos, err := n.NikkiRepository.FindPhotos(nikki.ID, nikki.NumberOfPhotos)
		nikki.Photos = photos
		return nikki, err
	}
	return nikki, nil
}

func (n *NikkiService) RegisterPhoto(NikkiID int, UserID int, Date int, PhotoID int, Photo string) {
	n.NikkiRepository.InsertPhoto(NikkiID, UserID, Date, PhotoID, Photo)
}

// func (n *NikkiService) GetAllPhotos() {
// 	n.NikkiRepository.FindAllPhotos()
// }
