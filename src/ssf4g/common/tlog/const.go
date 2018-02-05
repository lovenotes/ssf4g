package tlog

import (
	"log"
	"os"
	"time"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	LOG_LEVEL_ERROR = 1
	LOG_LEVEL_WARN  = 2
	LOG_LEVEL_INFO  = 3
	LOG_LEVEL_DEBUG = 4
)

const (
	CHECK_INTERVAL = 2 * time.Minute
	CHECK_EXPIRED  = 2 * time.Hour
)

var (
	_std_error = log.New(os.Stderr, "\033[0;31mERROR:\033[0m ", log.LstdFlags|log.Lshortfile)
	_std_info  = log.New(os.Stderr, "INFO : ", log.LstdFlags|log.Lshortfile)
	_std_warn  = log.New(os.Stderr, "\033[0;35mWARN :\033[0m ", log.LstdFlags|log.Lshortfile)
)
