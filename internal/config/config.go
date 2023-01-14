package config

import (
	"os"
	"strconv"
	"time"

	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
)

type appConfig struct {
	HTTPInfo  *httpInfo
	MySQLInfo *mysqlInfo
}

type httpInfo struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type mysqlInfo struct {
	User     string
	PassWord string
	Addr     string
	DBName   string
}

func LoadConfig() (*appConfig, error) {
	httpInfo, err := createHTTPInfo()
	if err != nil {
		return nil, err
	}

	mysqlInfo, err := createMySQLInfo()
	if err != nil {
		return nil, err
	}

	return &appConfig{
		HTTPInfo:  httpInfo,
		MySQLInfo: mysqlInfo,
	}, nil
}

func createHTTPInfo() (*httpInfo, error) {
	port, b := os.LookupEnv("HTTP_PORT")
	if !b {
		return nil, errorutil.Errorf(errorutil.Unknown, "env of HTTP_PORT is empty")
	}

	rTimeout, err := strconv.Atoi(os.Getenv("HTTP_ReadTimeout"))
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	wTimeout, err := strconv.Atoi(os.Getenv("HTTP_WriteTimeout"))
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	return &httpInfo{
		Addr:         ":" + port,
		ReadTimeout:  time.Duration(rTimeout) * time.Second,
		WriteTimeout: time.Duration(wTimeout) * time.Second,
	}, nil
}

func createMySQLInfo() (*mysqlInfo, error) {
	user, b := os.LookupEnv("MYSQL_USER")
	if !b {
		return nil, errorutil.Errorf(errorutil.Unknown, "env of MYSQL_USER is empty")
	}

	password, b := os.LookupEnv("MYSQL_PASSWORD")
	if !b {
		return nil, errorutil.Errorf(errorutil.Unknown, "env of MYSQL_PASSWORD is empty")
	}

	addr, b := os.LookupEnv("MYSQL_ADDR")
	if !b {
		return nil, errorutil.Errorf(errorutil.Unknown, "env of MYSQL_ADDR is empty")
	}

	dbName, b := os.LookupEnv("MYSQL_DATABASE")
	if !b {
		return nil, errorutil.Errorf(errorutil.Unknown, "env of MYSQL_DATABASE is empty")
	}

	return &mysqlInfo{
		User:     user,
		PassWord: password,
		Addr:     addr,
		DBName:   dbName,
	}, nil
}
