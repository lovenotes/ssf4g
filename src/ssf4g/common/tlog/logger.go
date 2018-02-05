package tlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	_file *os.File

	_debug *log.Logger
	_info  *log.Logger
	_warn  *log.Logger
	_err   *log.Logger

	_level int

	_rotate *rotate

	_lock sync.RWMutex
}

type rotate struct {
	_size int64

	_expired  time.Duration
	_interval time.Duration
}

func NewLogger(logpath string, loglevel int) *Logger {
	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		_std_error.Fatalf("new logger (%s) err (%v).", logpath, err)
	}

	logger := &Logger{
		_file: file,

		_level: loglevel,

		_rotate: &rotate{
			_size:     GB,
			_expired:  time.Hour * 24 * 7,
			_interval: time.Hour,
		},
	}

	logger.setLogLevel(file, loglevel)

	go logger.loggerMonitor()

	return logger
}

// Func - 设置Logger级别
func (logger *Logger) setLogLevel(file *os.File, loglevel int) {
	switch {
	case loglevel >= LOG_LEVEL_DEBUG:
		logger._debug = log.New(file, "\033[0;36mDEBUG:\033[0m ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case loglevel >= LOG_LEVEL_INFO:
		logger._info = log.New(file, "INFO : ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case loglevel >= LOG_LEVEL_WARN:
		logger._warn = log.New(file, "\033[0;35mWARN :\033[0m ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case loglevel >= LOG_LEVEL_ERROR:
		logger._err = log.New(file, "\033[0;31mERROR:\033[0m ", log.LstdFlags|log.Lshortfile)
	}

	switch {
	case loglevel < LOG_LEVEL_ERROR:
		logger._err = nil
		fallthrough
	case loglevel < LOG_LEVEL_WARN:
		logger._warn = nil
		fallthrough
	case loglevel < LOG_LEVEL_INFO:
		logger._info = nil
		fallthrough
	case loglevel < LOG_LEVEL_DEBUG:
		logger._debug = nil
	}
}

// Func - 获取TLog文件大小
func (logger *Logger) getFileSize() int64 {
	logger._lock.RLock()
	defer logger._lock.RUnlock()

	fi, err := logger._file.Stat()
	if err != nil {
		_std_warn.Printf("get file size err (%v).\n", err)

		return 0
	}

	return fi.Size()
}

// Func - 截断并重命名超过Interval的Log文件
func (logger *Logger) trunc(filepath, ext string) {
	logger._lock.Lock()
	defer logger._lock.Unlock()

	err := logger._file.Close()

	if err != nil {
		_std_warn.Printf("trunc err (close %v).\n", err)

		return
	}

	err = os.Rename(filepath, filepath+ext)

	if err != nil {
		_std_warn.Printf("trunc err (rename %v).\n", err)
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		_std_warn.Printf("trunc err (open %v).\n", err)

		return
	}

	// 重置文件写入
	logger.setLogLevel(file, logger._level)

	logger._file = file
}

// Func - 生成Logg文件后缀
func suffix(t time.Time) string {
	year, month, day := t.Date()

	return "-" + fmt.Sprintf("%04d%02d%02d%02d", year, month, day, t.Hour())
}

// Func - 截断获得下一次指定时间段的时间
func toNextBound(duration time.Duration) time.Duration {
	return time.Now().Truncate(duration).Add(duration).Sub(time.Now())
}

// Func - Log监听处理函数
func (logger *Logger) loggerMonitor() error {
	interval := time.After(toNextBound(logger._rotate._interval))
	expired := time.After(CHECK_EXPIRED)

	// 按照文件大小分割文件后缀
	sizeExt := 1

	fn := filepath.Base(logger._file.Name())

	fp, err := filepath.Abs(logger._file.Name())

	if err != nil {
		_std_error.Fatalf("logger monitor err (%v).", err)
	}

	for {
		var size <-chan time.Time
		if toNextBound(logger._rotate._interval) != CHECK_INTERVAL {
			size = time.After(CHECK_INTERVAL)
		}
		select {
		case t := <-interval:
			// 自定义生成新的Logger文件
			interval = time.After(logger._rotate._interval)
			logger.trunc(fp, suffix(t))
			sizeExt = 1

			_std_info.Printf("logger monitor info (truncated by interval).\n")
		case <-expired:
			// 删除过期的Logger文件
			expired = time.After(CHECK_EXPIRED)

			err := filepath.Walk(filepath.Dir(fp),
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return nil
					}

					isLog := strings.Contains(info.Name(), fn)

					if time.Since(info.ModTime()) > logger._rotate._expired && isLog && info.IsDir() == false {
						if err := os.Remove(path); err != nil {
							return err
						}

						_std_info.Printf("logger monitor (%s) info (remove by expired).\n", path)
					}
					return nil
				})

			if err != nil {
				_std_error.Printf("logger monitor err (remove %v).\n", err)
			}
		case t := <-size:
			// 文件大小超过上限
			if logger.getFileSize() < logger._rotate._size {
				break
			}

			logger.trunc(fp, suffix(t)+"."+strconv.Itoa(sizeExt))

			sizeExt++

			_std_info.Printf("logger monitor info (trunc by size).\n")
		}
	}
}

// Func - 输出Debug日志
func (logger *Logger) Debug(format string, v ...interface{}) {
	logger._lock.RLock()
	defer logger._lock.RUnlock()

	if logger._debug != nil {
		logger._debug.Output(3, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}

// Func - 输出Info日志
func (logger *Logger) Info(format string, v ...interface{}) {
	logger._lock.RLock()
	defer logger._lock.RUnlock()

	if logger._info != nil {
		logger._info.Output(3, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}

// Func - 输出Warn日志
func (logger *Logger) Warn(format string, v ...interface{}) {
	_std_warn.Output(3, fmt.Sprintln(fmt.Sprintf(format, v...)))

	logger._lock.RLock()
	defer logger._lock.RUnlock()

	if logger._warn != nil {
		logger._warn.Output(3, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}

// Func - 输出Error日志
func (logger *Logger) Error(format string, v ...interface{}) {
	_std_error.Output(3, fmt.Sprintln(fmt.Sprintf(format, v...)))

	logger._lock.RLock()
	defer logger._lock.RUnlock()

	if logger._err != nil {
		logger._err.Output(3, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}
