//
// stat.Stat describes data in /proc/<pid>/stat.
//
// Use stat.New() to create a new stat.Stat object
// from data in a path.
//
package stat

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/chair300/procfs/util"
)

//
// Stat is a data structure that maps to /proc/<pid>/stat.
//
type Stat struct {
	Pid                 int           // process id
	Comm                string        // filename of the executable
	State               string        // state (R is running, S is sleeping, D is sleeping in an uninterruptible wait, Z is zombie, T is traced or stopped)
	Ppid                int           // process id of the parent process
	Pgrp                int           // pgrp of the process
	Session             int           // session id
	TtyNr               int           // tty the process uses
	Tpgid               int           // pgrp of the tty
	Flags               int64         // task flags
	Minflt              int64         // number of minor faults
	Cminflt             int64         // number of minor faults with child's
	Majflt              int64         // number of major faults
	Cmajflt             int64         // number of major faults with child's
	Utime               time.Duration // user mode in nanoseconds
	Stime               time.Duration // kernel mode in nanoseconds
	Cutime              time.Duration // user mode in nanoseconds with child's
	Cstime              time.Duration // kernel mode in nanoseconds with child's
	Priority            int64         // priority level
	Nice                int64         // nice level
	NumThreads          int64         // number of threads
	Itrealvalue         int64         // (obsolete, always 0)
	Starttime           time.Time     // time the process started after system boot
	Vsize               int64         // virtual memory size
	Rss                 int64         // resident set memory size
	Rlim                uint64        // current limit in bytes on the rss
	Startcode           int64         // address above which program text can run
	Endcode             int64         // address below which program text can run
	Startstack          int64         // address of the start of the main process stack
	Kstkesp             int64         // current value of ESP
	Kstkeip             int64         // current value of EIP
	Signal              int64         // bitmap of pending signals
	Blocked             int64         // bitmap of blocked signals
	Sigignore           int64         // bitmap of ignored signals
	Sigcatch            int64         // bitmap of catched signals
	Wchan               uint64        // address where process went to sleep
	Nswap               int64         // (place holder)
	Cnswap              int64         // (place holder)
	ExitSignal          int           // signal to send to parent thread on exit
	Processor           int           // which CPU the task is scheduled on
	RtPriority          int64         // realtime priority
	Policy              int64         // scheduling policy (man sched_setscheduler)
	DelayacctBlkioTicks int64         // time spent waiting for block IO
}
var Desc  = map[string]label{struct {
	"Pid"  	: label{
          		unit: int,
				description:  " process id",
			},
	"Comm" 	: label{
          		unit: string,
				description:  "filename of the executable",
			},
	"State" 	: label{
          		unit: string,
				description:  "state (R is running, S is sleeping, D is sleeping in an uninterruptible wait, Z is zombie, T is traced or stopped)",
			},
	"Ppid"    : label{
          		unit: int,
				description:  "process id of the parent process",
			},
	"Pgrp"    : label{
          		unit: int,
				description:  "pgrp of the process",
			},
	"Session" : label{
          		unit: int,
				description:  "session id",
			},
	"TtyNr"  : label{
          		unit: int,
				description:  "tty the process uses",
			},
	"Tpgid"  : label{
          		unit: int,
				description:  "pgrp of the tty",
			},
	"Flags"   : label{
          		unit: int64,
				description:  "task flags",
			},
	"Minflt"  : label{
          		unit: int64,
				description:  "number of minor faults",
			},
	"Cminflt" : label{
          		unit: int64,
				description:  "number of minor faults with child's",
			},
	"Majflt"  : label{
          		unit: int64,
				description:  "number of major faults",
			},
	"Cmajflt" : label{
          		unit: int64,
				description:  "number of major faults with child's",
			},
	"Utime"   : label{
          		unit: time.Duration,
				description:  "user mode in nanoseconds",
			},
	"Stime"   : label{
          		unit:time.Duration,
				description:  "kernel mode in nanoseconds",
			},
	"Cutime"  : label{
          		unit:  time.Duration ,
				description:  "user mode in nanoseconds with child's",
			},
	"Cstime"  : label{
          		unit: time.Duration ,
				description:  "kernel mode in nanoseconds with child's",
			},
	"Priority": label{
          		unit: int64,
				description:  "priority level",
			},
	"Nice"    : label{
          		unit: int64,
				description:  "nice level",
			},
	"NumThreads" : label{
          		unit: int64,
				description:  "number of threads",
			},
	"Itrealvalue": label{
          		unit: int64,
				description:  "(obsolete, always 0)",
			},
	"Starttime"  : label{
          		unit: int64,
				description:  "time the process started after system boot",
			},
	"Vsize"  : label{
          		unit: int64,
				description:  "virtual memory size",
			},
	"Rss"    : label{
          		unit: int64,
				description:  "resident set memory size",
			},
	"Rlim"   : label{
          		unit: int64,
				description:  "current limit in bytes on the rss",
			},
	"Startcode" : label{
          		unit: int64,
				description:  "address above which program text can run",
			},
	"Endcode"  : label{
          		unit: int64,
				description:  "address below which program text can run",
			},
	"Startstack": label{
          		unit: int64,
				description:  "address of the start of the main process stack",
			},
	"Kstkesp"  : label{
          		unit: int64,
				description:  "current value of ESP",
			},
	"Kstkeip"  : label{
          		unit: int64,
				description:  "current value of EIP",
			},
	"Signal"   : label{
          		unit: int64,
				description:  "bitmap of pending signals",
			},
	"Blocked"  : label{
          		unit: int64,
				description:  " bitmap of blocked signals",
			},
	"Sigignore" : label{
          		unit: int64,
				description:  "bitmap of ignored signals",
			},
	"Sigcatch"  : label{
          		unit: int64,
				description:  "bitmap of catched signals",
			},
	"Wchan"  : label{
          		unit: int64,
				description:  "address where process went to sleep)",
			},
	"Nswap"  : label{
          		unit: int64,
				description:  "(place holder)",
			},
	"Cnswap"  : label{
          		unit: int64,
				description:  "(place holder)",
			},
	"ExitSignal"  : label{
          		unit: int,
				description:  "signal to send to parent thread on exit",
			},
	"Processor" : label{
          		unit: int,
				description:  " which CPU the task is scheduled on",
			},
	"RtPriority" : label{
          		unit: int64,
				description:  "realtime priority",
			},
	"Policy" : label{
          		unit: int64,
				description:  "scheduling policy (man sched_setscheduler)",
			},
	"DelayacctBlkioTicks": label{
          		unit: int64,
				description:  "time spent waiting for block IO",
			},
	}
}

//
// stat.New creates a new /proc/<pid>/stat from a path.
//
// An error is returned if the data is malformed, or the path does not exist.
//
func New(path string) (*Stat, error) {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buf), " ")
	stat := &Stat{}
	err = util.ParseStringsIntoStruct(stat, lines)
	return stat, err
}
