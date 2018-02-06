package utility

import (
	"os"
	"strconv"

	"ssf4g/common/tlog"
)

func GenPIDFile() {
	pid := os.Getpid()

	pidPath := os.Args[0] + ".pid"
	pidFile, err := os.OpenFile(pidPath, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		tlog.Error("gen pid file (%s) err (open %v).", pidPath, err)

		os.Exit(-1)
	}

	_, err = pidFile.WriteString(strconv.Itoa(pid))

	if err != nil {
		tlog.Error("gen pid file (%s) err (write %v).", pidPath, err)

		os.Exit(-1)
	} else {
		tlog.Info("gen pid file (%s) info (generated).", pidPath)
	}
}

func RemovePIDFile() {
	pidPath := os.Args[0] + ".pid"
	os.Remove(pidPath)

	tlog.Info("remove pid file (%s) info (removed).", pidPath)
}
