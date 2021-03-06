package cli

import (
	"fmt"
	"qiniu/api.v6/auth/digest"
	"qiniu/log"
	"qshell"
	"time"
)

func Sync(cmd string, params ...string) {
	if len(params) == 3 || len(params) == 4 {
		srcResUrl := params[0]
		bucket := params[1]
		key := params[2]
		upHostIp := ""
		if len(params) == 4 {
			upHostIp = params[3]
		}

		gErr := accountS.Get()
		if gErr != nil {
			log.Error(gErr)
			return
		}

		mac := digest.Mac{
			accountS.AccessKey,
			[]byte(accountS.SecretKey),
		}

		//sync
		tStart := time.Now()
		hash, sErr := qshell.Sync(&mac, srcResUrl, bucket, key, upHostIp)
		if sErr != nil {
			log.Error(sErr)
			return
		}

		fmt.Println(fmt.Sprintf("Sync %s => %s:%s (%s) Success, Duration: %s!",
			srcResUrl, bucket, key, hash, time.Since(tStart)))
	} else {
		CmdHelp(cmd)
	}
}
