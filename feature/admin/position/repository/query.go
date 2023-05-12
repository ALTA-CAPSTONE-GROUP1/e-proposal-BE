package repository

import (
	"fmt"

	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin"
	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin/position"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type positionModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) position.Repository {
	return &positionModel{
		db: db,
	}
}

func (pm *positionModel) InsertPosition(position position.Core) error {
	newPosition := admin.Position{
		Name: position.Name,
		Tag:  position.Tag,
	}

	tx := pm.db.Create(&newPosition)
	if tx.Error != nil {
		log.Errorf("error on insert data %s with tag %s to position table in db", position.Name, position.Tag)
		return tx.Error
	}

	return nil
}

func (pm *positionModel) GetPositions(limit int, offset int, search string) ([]position.Core, error) {
	var (
		positions   []position.Core
		DBpositions []admin.Position
	)

	if search != "" {
		searchConds := ("%" + search + "%")

		tx := pm.db.Limit(limit).Offset(offset).Where("name LIKE ? OR tag LIKE ?", searchConds, searchConds).Find(&DBpositions)
		if tx.Error != nil {
			log.Error("get positions query error with search condition")
			return nil, tx.Error
		}

		for _, dbPos := range DBpositions {
			corePos := position.Core{
				Name: dbPos.Name,
				Tag:  dbPos.Tag,
			}
			positions = append(positions, corePos)
		}

		return positions, nil
	}

	tx := pm.db.Limit(limit).Offset(offset).Find(&DBpositions)
	if tx.Error != nil {
		log.Error("get posititons query error without search condition")
		return nil, tx.Error
	}

	for _, dbPos := range DBpositions {
		corePos := position.Core{
			Name: dbPos.Name,
			Tag:  dbPos.Tag,
		}
		positions = append(positions, corePos)
	}

	return positions, nil
}

func (pm *positionModel) DeletePosition(position string, tag string) error {
	var count int64
	tx := pm.db.Model(&admin.Position{}).Where(&admin.Position{Name: position, Tag: tag}).Count(&count)
	if tx.Error != nil {
		log.Error("count position query error")
		return fmt.Errorf("count position query error: %w", tx.Error)
	}

	if count == 0 {
		log.Warn("no position data found for deletion")
		return fmt.Errorf("no position data found for deletion with name %s and tag %s", position, tag)
	}

	tx = pm.db.Delete(&admin.Position{Name: position, Tag: tag})
	if tx.Error != nil {
		log.Error("delete position query error")
		return fmt.Errorf("delete position query error: %w", tx.Error)
	}
	return nil
}
