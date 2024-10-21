package core

import (
	"sync"

	var_db "github.com/flabio/safe_var_db"
	"github.com/safe_msvc_user/insfractruture/database"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
	"github.com/safe_msvc_user/usecase/dto"
)

func NewSchoolRepository() uicore.UISchoolCore {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}

func (c *OpenConnection) GetSchoolFindAll(begin int) ([]dto.SchoolResponseDTO, int64, error) {
	var schoolEntities []dto.SchoolResponseDTO
	var countSchool int64
	c.mux.Lock()
	query := c.connection.Table("schools").Offset(begin).Limit(utils.LIMIT).Order(var_db.DB_ORDER_DESC).Find(&schoolEntities)
	c.connection.Table("schools").Count(&countSchool)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return schoolEntities, countSchool, query.Error
}
func (c *OpenConnection) GetSchoolFindById(id uint) (entities.School, error) {
	var school entities.School
	c.mux.Lock()
	result := c.connection.Where(var_db.DB_EQUAL_ID, id).Find(&school)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, result.Error
}
func (c *OpenConnection) GetSchoolFindByEmail(id uint, email string) (entities.School, error) {
	var school entities.School
	c.mux.Lock()
	query := c.connection.Where(var_db.DB_EQUAL_EMAIL_ID, email)
	if id > 0 {
		query = query.Where(var_db.DB_DIFF_ID, id)
	}
	query = query.Find(&school)

	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, query.Error
}
func (c *OpenConnection) GetSchoolFindByProviderNumber(id uint, providerNumber string) (entities.School, error) {
	var school entities.School
	c.mux.Lock()
	query := c.connection.Where("provider_number=?", providerNumber)
	if id > 0 {
		query = query.Where(var_db.DB_DIFF_ID, id)
	}
	query = query.Find(&school)

	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, query.Error
}

func (c *OpenConnection) CreateSchool(school entities.School) (entities.School, error) {
	c.mux.Lock()
	err := c.connection.Create(&school).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, err
}
func (c *OpenConnection) UpdateSchool(id uint, school entities.School) (entities.School, error) {
	c.mux.Lock()
	err := c.connection.Where(var_db.DB_EQUAL_ID, id).Updates(&school).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, err
}

func (c *OpenConnection) DeleteSchool(id uint) (bool, error) {
	c.mux.Lock()
	var school entities.School
	err := c.connection.Where(var_db.DB_EQUAL_ID, id).Delete(&school).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return err == nil, err
}
