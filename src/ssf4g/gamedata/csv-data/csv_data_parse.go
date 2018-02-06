package gamedata

import (
	"encoding/csv"
	"os"
	"path"
	"strings"

	"ssf4g/common/tlog"
)

func setData(tblname string, rowname string, fieldname string, value string) {
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

func dataParse(file *os.File) error {
	// csv 读取器
	csv_reader := csv.NewReader(file)

	records, err := csv_reader.ReadAll()

	if err != nil {
		tlog.Error("data parse (%s) err (read all %v).", file.Name(), err)

		return err
	}

	// 是否为空文件
	if len(records) == 0 {
		tlog.Error("data parse (%s) err (file empty).", file.Name())

		return nil
	}

	// 处理表名
	fi, err := file.Stat()

	if err != nil {
		tlog.Error("data parse (%s) err (stat %v).", file.Name(), err)

		return err
	}

	tblname := strings.TrimSuffix(fi.Name(), path.Ext(file.Name()))

	// 记录数据, 第一行为表头，因此从第二行开始
	for line := 1; line < len(records); line++ {
		for field := 1; field < len(records[line]); field++ { // 每条记录的第一个字段作为行索引
			setData(tblname, records[line][0], records[0][field], records[line][field])
		}
	}

	return nil
}
