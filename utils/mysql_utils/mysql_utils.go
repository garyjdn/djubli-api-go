package mysql_utils

import (
	"github.com/garyjdn/djubli-api-go/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NotFoundError("no record matching given id")
		}
		return errors.InternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("invalid data")
	}
	return errors.InternalServerError("error processing request")
}
