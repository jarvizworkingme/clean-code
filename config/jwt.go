package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Jwt struct {
	Key     string
	Expired int
}

func NewJwtConfig() Jwt {
	exp, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		logrus.Errorln("LOAD ENV JWT_EXPIRED FAILED ", err.Error())
		exp = 60 * 24
	}

	return Jwt{
		Key:     os.Getenv("JWT_KEY"),
		Expired: exp,
	}

}
