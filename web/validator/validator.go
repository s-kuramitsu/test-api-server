package validator

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/s-kuramitsu/my-api/entity"
)

// Validate email regexp
var emailValidator = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

// Validate name regexp
var nameValidator = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+$")

// Validate add user request body
func AddUserValidator(context *gin.Context) *entity.User {
	user := &entity.User{}
	context.BindJSON(user)
	if user.ID != 0 {
		return nil
	}
	if !emailValidator.Match([]byte(user.Email)) {
		return nil
	}
	if !nameValidator.Match([]byte(user.Name)) {
		return nil
	}
	if !user.CreatedAt.IsZero() || !user.UpdatedAt.IsZero() || user.DeletedAt != nil {
		return nil
	}
	return user
}
