/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Monday 14 October 2019 - 17:54:41
** @Filename:				Mutex.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Friday 18 October 2019 - 13:34:20
*******************************************************************************/

package		mut

import		"sync"
import		"github.com/microgolang/logs"

var		MUTEXES = map[string]MutArr{}
var		LOGS = false

type	MutArr struct {
	Mut		*sync.Mutex
	uid		string
}

func	SetLogs(shouldSet bool) {
	LOGS = shouldSet
}

func	get(uid string) MutArr {
	_, ok := MUTEXES[uid]
	if (!ok) {
		MUTEXES[uid] = MutArr{
			Mut: &sync.Mutex{},
			uid: uid,
		}
	}
	return (MUTEXES[uid])
}

func	Lock(uid string) {
	Element := get(uid)
	Element.Mut.Lock()
	if (LOGS) {
		logs.Info(`Lock ` + uid)
	}
}

func	Unlock(uid string) {
    defer func() {
        if r := recover(); r != nil {
            return
        }
    }()
	Element := get(uid)
	Element.Mut.Unlock()
	if (LOGS) {
		logs.Info(`Unlock ` + uid)
	}
}
