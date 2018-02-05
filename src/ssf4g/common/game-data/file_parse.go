package gamedata

import (
	"encoding/csv"
	"os"
	"path"
	"strings"

	"ssf4g/libs/tlog"
)

//----------------------- parse & load a game data file into dictionary ------------------------
func fileParse(file *os.File) error {
	// csv 读取器
	csv_reader := csv.NewReader(file)

	records, err := csv_reader.ReadAll()

	if err != nil {
		tlog.Error("file parse (%s) err (read all %v).", file.Name(), err)

		return err
	}

	// 是否为空文件
	if len(records) == 0 {
		logger.GetNLog().Error("load csv file (%s) err (file empty).", file.Name())

		return
	}

	// 处理表名
	fi, err := file.Stat()

	if err != nil {
		logger.GetNLog().Error("stat the file (%s) err (%v).", file.Name(), err)

		return
	}

	tblname := strings.TrimSuffix(fi.Name(), path.Ext(file.Name()))

	// 记录数据, 第一行为表头，因此从第二行开始
	for line := 1; line < len(records); line++ {
		for field := 1; field < len(records[line]); field++ { // 每条记录的第一个字段作为行索引
			_set(tblname, records[line][0], records[0][field], records[line][field])
		}
	}

	return nil
}
