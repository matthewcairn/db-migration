package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gorm "github.com/jinzhu/gorm"
	"transfer/pkg/model"
)

func ParseIDFromUri(c *gin.Context) *uuid.UUID {
	tID := model.UriParse{}
	if err := c.ShouldBindUri(&tID); err != nil {
		_ = c.Error(err)
		return nil
	}
	if len(tID.ID) == 0 {
		_ = c.Error(fmt.Errorf("error: Empty when parse ID from URI"))
		return nil
	}
	if id, err := uuid.Parse(tID.ID[0]); err != nil {
		_ = c.Error(err)
		return nil
	} else {
		return &id
	}
}

func IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func GetKeyOrNil(jsonInput map[string]interface{}, value string) interface{} {
	if val, ok := jsonInput[value]; ok {
		return val
	} else {
		return nil
	}

}
