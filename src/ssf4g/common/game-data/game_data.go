package gamedata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"bosslove/common/config"
	"bosslove/common/logger"
)

var (
	_lock   sync.RWMutex
	_tables map[string]*Table
)

const (
	GAME_DATA_OK = iota

	GAME_DATA_FORMAT_ERROR
	GAME_DATA_NOT_EXISTS
)

//------------------------------ info for a level ------------------------------
type Record struct {
	Fields map[string]string
}

//------------------------------ Numerical Table for a object ------------------------------
type Table struct {
	Records map[string]*Record
}

func InitGameData(gamedatadir string) {
	Reload(gamedatadir)
}

//------------------------------ Reload *.csv ------------------------------
func Reload(gamedatadir string) {
	_lock.Lock()
	defer _lock.Unlock()

	_tables = make(map[string]*Table)

	pattern := os.Getenv("GOPATH") + "/src/gamedata/data/*.csv"

	if gamedatadir != "" {
		pattern = gamedatadir + "/*.csv"
	}

	logger.GetNLog().Info("loading gamedata from %s", pattern)
	files, err := filepath.Glob(pattern)

	if err != nil {
		logger.GetNLog().Error("err file path glob (%v)", err)
		panic(err)
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			logger.GetNLog().Error("error opening file (%v)", err)
			continue
		}

		parse(file)
		file.Close()
	}

	logger.GetNLog().Info("%v CSV(s) Loaded", len(_tables))
}

//------------------------------ Set Field value ------------------------------
func _set(tblname string, rowname string, fieldname string, value string) {
	tbl := _tables[tblname]

	if tbl == nil {
		tbl = &Table{}
		tbl.Records = make(map[string]*Record)
		_tables[tblname] = tbl
	}

	rec := tbl.Records[rowname]
	if rec == nil {
		rec = &Record{}
		rec.Fields = make(map[string]string)
		tbl.Records[rowname] = rec
	}

	rec.Fields[fieldname] = value
}

//------------------------------ Get Field value ------------------------------
func _get(tblname string, rowname string, fieldname string) (string, int) {
	_lock.RLock()
	defer _lock.RUnlock()

	tbl, ok := _tables[tblname]
	if !ok {
		logger.GetNLog().Error(fmt.Sprint("table ", tblname, " not exists!"))
		return "", GAME_DATA_NOT_EXISTS
	}

	rec, ok := tbl.Records[rowname]
	if !ok {
		logger.GetNLog().Error(fmt.Sprint("table ", tblname, " row ", rowname, " not exists!"))
		return "", GAME_DATA_NOT_EXISTS
	}

	value, ok := rec.Fields[fieldname]
	if !ok {
		logger.GetNLog().Error(fmt.Sprint("table ", tblname, " field ", fieldname, " not exists!"))
		return "", GAME_DATA_NOT_EXISTS
	}
	return value, GAME_DATA_OK
}

// ------------------------------ Get Table Info ------------------------------
func GetTable(tblname string) (map[string]*Record, int) {
	_lock.RLock()
	defer _lock.RUnlock()

	tbl, ok := _tables[tblname]
	if !ok {
		logger.GetNLog().Error(fmt.Sprint("table ", tblname, " not exists!"))
		return nil, GAME_DATA_NOT_EXISTS
	}

	return tbl.Records, GAME_DATA_OK
}

//------------------------------ Get Field value as Integer ------------------------------
func GetInt(tblname string, rowname string, fieldname string) (int32, int) {
	val, ret := _get(tblname, rowname, fieldname)
	if ret != GAME_DATA_OK {
		return 0, ret
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		logger.GetNLog().Error(fmt.Sprintf("parse integer from gamedata (%v %v %v) err (%v).", tblname, rowname, fieldname, err))
		return 0, GAME_DATA_FORMAT_ERROR
	}

	return int32(v), GAME_DATA_OK
}

//------------------------------ Get Field value as Float ------------------------------
func GetFloat(tblname string, rowname string, fieldname string) (float64, int) {
	val, ret := _get(tblname, rowname, fieldname)
	if ret != GAME_DATA_OK {
		return 0.0, ret
	}

	/*
		if val == "" {
			return 0.0, GAME_DATA_OK
		}
	*/

	f, err := strconv.ParseFloat(val, 32)
	if err != nil {
		logger.GetNLog().Error(fmt.Sprintf("parse float from gamedata (%v %v %v) err (%v).", tblname, rowname, fieldname, err))
		return 0.0, GAME_DATA_FORMAT_ERROR
	}

	return f, GAME_DATA_OK
}

//------------------------------ Get Field value as string ------------------------------
func GetString(tblname string, rowname string, fieldname string) (string, int) {
	return _get(tblname, rowname, fieldname)
}

//------------------------------ Get Row Count ------------------------------
func Count(tblname string) int32 {
	tbl := _tables[tblname]

	if tbl == nil {
		return 0
	}

	return int32(len(tbl.Records))
}

//------------------------------ Test Field Exists ------------------------------
func IsFieldExists(tblname string, fieldname string) bool {
	_lock.RLock()
	defer _lock.RUnlock()

	tbl := _tables[tblname]

	if tbl == nil {
		return false
	}

	key := ""
	// get one record key
	for k := range tbl.Records {
		key = k
		break
	}

	rec, ok := tbl.Records[key]
	if !ok {
		return false
	}

	_, ok = rec.Fields[fieldname]
	if !ok {
		return false
	}

	return true
}

//------------------------------ Load JSON From GameData Directory ------------------------------
func LoadJSON(filename string) ([]byte, error) {
	prefix := os.Getenv("GOPATH") + "/src/gamedata/data"
	config := config.Get()
	if config["gamedata_dir"] != "" {
		prefix = config["gamedata_dir"]
	}

	path := prefix + "/" + filename
	return ioutil.ReadFile(path)
}
