package cli

import (
	"qiniu/log"
	"qshell"
	"strconv"
)

func CheckQrsync(cmd string, params ...string) {
	if len(params) == 3 || len(params) == 4 {
		dirCacheResultFile := params[0]
		listBucketResultFile := params[1]
		ignoreLocalDir, err := strconv.ParseBool(params[2])
		if err != nil {
			log.Errorf("Invalid value `%s' for argument <IgnoreLocalDir>", params[2])
			return
		}
		prefix := ""
		if len(params) == 4 {
			prefix = params[3]
		}
		qshell.CheckQrsync(dirCacheResultFile, listBucketResultFile, ignoreLocalDir, prefix)
	} else {
		CmdHelp(cmd)
	}
}
