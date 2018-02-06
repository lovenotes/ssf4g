package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io/ioutil"
)

// Func - 压缩数据为Gzip格式
func DataToGzip(data []byte) ([]byte, error) {
	retData := make([]byte, 0)

	if len(data) <= GZIP_BYTE_LIMIT {
		retData = append(retData, '0')
		retData = append(retData, data...)

		return retData, nil
	}

	buf := new(bytes.Buffer)

	writer := gzip.NewWriter(buf)

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

	err = writer.Close()

	if err != nil {
		return nil, err
	}

	retData = append(retData, '1')
	retData = append(retData, buf.Bytes()...)

	return retData, nil
}

// Func - 解压Gzip格式数据
func GzipToData(data []byte) ([]byte, error) {
	if data[0] == '0' {
		return data[1:], nil
	}

	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, data[1:])

	reader, err := gzip.NewReader(buf)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	udata, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	return udata, nil
}
