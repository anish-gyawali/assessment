package service

import (
	"log"

	"github.com/anish-gyawali/assessment/server/dto"
	"github.com/anish-gyawali/assessment/server/entity"
	"github.com/anish-gyawali/assessment/server/repository"
	"github.com/mashingan/smapping"
)

//DataService is a ....
type DataService interface {
	Insert(b dto.DataCreateDTO) entity.Data
	Update(b dto.DataUpdateDTO) entity.Data
	Delete(b entity.Data)
	All() []entity.Data
	FindByID(dataID int64) entity.Data
}

type dataService struct {
	dataRepository repository.DataRepository
}

//NewDataService .....
func NewDataService(dataRepo repository.DataRepository) DataService {
	return &dataService{
		dataRepository: dataRepo,
	}
}

func (service *dataService) Insert(b dto.DataCreateDTO) entity.Data {
	data := entity.Data{}
	err := smapping.FillStruct(&data, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.dataRepository.InsertData(data)
	return res
}

func (service *dataService) Update(b dto.DataUpdateDTO) entity.Data {
	data := entity.Data{}
	err := smapping.FillStruct(&data, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.dataRepository.UpdateData(data)
	return res
}

func (service *dataService) Delete(b entity.Data) {
	service.dataRepository.DeleteData(b)
}

func (service *dataService) All() []entity.Data {
	return service.dataRepository.AllData()
}

func (service *dataService) FindByID(dataID int64) entity.Data {
	return service.dataRepository.FindDataByID(dataID)
}
