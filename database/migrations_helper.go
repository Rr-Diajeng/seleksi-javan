package database

import (
	"fmt"

	"gorm.io/gorm"
)

func DefineEnums(db *gorm.DB) error {
	enumStatements := map[string]string{
		"status": "('pending', 'completed', 'in_progress')",
	}

	for typename, statement := range enumStatements {
		if !IsEnumExists(db, typename) {
			createStmt := fmt.Sprintf("CREATE TYPE %s AS ENUM %s", typename, statement)
			if err := db.Exec(createStmt).Error; err != nil {
				return err
			}
		}
	}

	return nil

}

func IsEnumExists(db *gorm.DB, enumName string) bool {
	var exists bool
	db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = ?)", enumName).Scan(&exists)

	return exists
}
