package compress

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"

	"github.com/golang/snappy"
)

// Func - 压缩数据为Snappy格式
func DataToSnappy(data []byte) ([]byte, error) {
	retData := make([]byte, 0)

	if len(data) <= SNAPPY_BYTE_LIMIT {
		retData = append(retData, '0')
		retData = append(retData, data...)

		return retData, nil
	}

	buf := new(bytes.Buffer)

	writer := snappy.NewBufferedWriter(buf)
	defer writer.Close()

	len, err := writer.Write(data)

	if err != nil {
		return nil, err
	} else if len == 0 {
		return nil, nil
	}

	err = writer.Flush()

	if err != nil {
		return nil, err
	}

	retData = append(retData, '1')
	retData = append(retData, buf.Bytes()...)

	return retData, nil
}

// Func - 解压Snappy格式数据
func SnappyToData(data []byte) ([]byte, error) {
	if data[0] == '0' {
		return data[1:], nil
	}

	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, data[1:])

	reader := snappy.NewReader(buf)

	udata, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	return udata, nil
}
