package csvdata

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"ssf4g/common/tlog"
)

var (
	_lock sync.RWMutex

	_tables map[string]*Table
)

const (
	CSV_DATA_OK = iota

	CSV_DATA_FORMAT_ERROR
	CSV_DATA_NOT_EXISTS
)

type Record struct {
	Fields map[string]string
}

type Table struct {
	Records map[string]*Record
}

func InitCsvData(csvpath string) {
	ReloadCsvData(csvpath)
}

func ReloadCsvData(csvpath string) {
	_lock.Lock()
	defer _lock.Unlock()

	_tables = make(map[string]*Table)

	pattern := os.Getenv("GOPATH") + "/src/data/*.csv"

	if csvpath != "" {
		pattern = csvpath + "/*.csv"
	}

	tlog.Info("reload csv data from %s.", pattern)

	files, err := filepath.Glob(pattern)

	if err != nil {
		tlog.Error("reload csv data (%s) err (glob %v).", pattern, err)

		return
	}

	for _, f := range files {
		file, err := os.Open(f)

		if err != nil {
			tlog.Error("reload csv data (%s) err (open %v).", f, err)

			continue
		}

		dataParse(file)

		file.Close()
	}

	tlog.Info("reload csv data %v CSV(s) loaded.", len(_tables))
}

func getData(tblname string, rowname string, fieldname string) (string, int) {
	_lock.RLock()
	defer _lock.RUnlock()

	table, ret := _tables[tblname]

	if ret == false {
		tlog.Error("get (%s) err (table not exists).", tblname)

		return "", CSV_DATA_NOT_EXISTS
	}

	record, ret := table.Records[rowname]

	if ret == false {
		tlog.Error("get (%s, %s) err (row not exists).", tblname, rowname)

		return "", CSV_DATA_NOT_EXISTS
	}

	value, ret := record.Fields[fieldname]

	if ret == false {
		tlog.Error("get (%s, %s) err (field not exists).", tblname, fieldname)

		return "", CSV_DATA_NOT_EXISTS
	}

	return value, CSV_DATA_OK
}

func GetTable(tblname string) (map[string]*Record, int) {
	_lock.RLock()
	defer _lock.RUnlock()

	tbl, ret := _tables[tblname]

	if ret == false {
		tlog.Error("get table (%s) err (table not exists).", tblname)

		return nil, CSV_DATA_NOT_EXISTS
	}

	return tbl.Records, CSV_DATA_OK
}

func GetInt(tblname string, rowname string, fieldname string) (int32, int) {
	val, ret := getData(tblname, rowname, fieldname)
	if ret != CSV_DATA_OK {
		return 0, ret
	}

	v, err := strconv.Atoi(val)

	if err != nil {
		tlog.Error("get int (%s, %s, %s) err (convert %v).", tblname, rowname, fieldname, err)

		return 0, CSV_DATA_FORMAT_ERROR
	}

	return int32(v), CSV_DATA_OK
}

func GetFloat(tblname string, rowname string, fieldname string) (float64, int) {
	val, ret := getData(tblname, rowname, fieldname)

	if ret != CSV_DATA_OK {
		return 0.0, ret
	}

	f, err := strconv.ParseFloat(val, 32)

	if err != nil {
		tlog.Error("get float (%s, %s, %s) err (convert %v).", tblname, rowname, fieldname, err)

		return 0.0, CSV_DATA_FORMAT_ERROR
	}

	return f, CSV_DATA_OK
}

func GetString(tblname string, rowname string, fieldname string) (string, int) {
	return getData(tblname, rowname, fieldname)
}

func Count(tblname string) int32 {
	table := _tables[tblname]

	if table == nil {
		return 0
	}

	return int32(len(table.Records))
}

func IsFieldExists(tblname string, fieldname string) bool {
	_lock.RLock()
	defer _lock.RUnlock()

	table := _tables[tblname]

	if table == nil {
		return false
	}

	key := ""

	// get one record key
	for k := range table.Records {
		key = k

		break
	}

	rec, ret := table.Records[key]

	if ret == false {
		return false
	}

	_, ret = rec.Fields[fieldname]

	if ret == false {
		return false
	}

	return true
}
