package repository

import (
	"github.com/anish-gyawali/assessment/server/entity"
	"gorm.io/gorm"
)

//DataRepository is a ....
type DataRepository interface {
	InsertData(b entity.Data) entity.Data
	UpdateData(b entity.Data) entity.Data
	DeleteData(b entity.Data)
	AllData() []entity.Data
	FindDataByID(dataID int64) entity.Data
}

type dataConnection struct {
	connection *gorm.DB
}

//NewDataRepository creates an instance DataRepository
func NewDataRepository(dbConn *gorm.DB) DataRepository {
	return &dataConnection{
		connection: dbConn,
	}
}

//Insert data
func (db *dataConnection) InsertData(b entity.Data) entity.Data {
	db.connection.Save(&b)
	return b
}

//Update data
func (db *dataConnection) UpdateData(b entity.Data) entity.Data {
	db.connection.Save(&b)
	return b
}

//Delete data
func (db *dataConnection) DeleteData(b entity.Data) {
	db.connection.Delete(&b)
}

//FInd data by Id
func (db *dataConnection) FindDataByID(dataID int64) entity.Data {
	var data entity.Data
	db.connection.Preload("Datas").Find(&data, dataID)
	return data
}

//Get all data
func (db *dataConnection) AllData() []entity.Data {
	var datas []entity.Data
	db.connection.Preload("Datas").Find(&datas)
	return datas
}
