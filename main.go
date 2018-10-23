/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2018 Great Lakes SAN, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package procfs

import (
     "fmt"
     "os"
     "strconv"
	"github.com/chair300/procfs"
)


func main() {
     argsWithoutProg := os.Args[1:]
     fmt.Println(argsWithoutProg)
     // fetch all processes from /proc
     // returns a map of pid -> *Process
//     procs, err := procfs.Processes(false);
     pid,err := strconv.Atoi(os.Args[1])
     procs, err := procfs.NewProcess( pid, false)
     //fmt.Println(procs)
     if err != nil {}
     k := procs.Pid
     v := procs

     fmt.Println("pid:", k)
     fmt.Println("Command:", v.Cmd)
     s, err := v.Statm()
     if err == nil {
          fmt.Println("size:", s.Size)
     }
     fmt.Println("size:", s.Size)
     v.GetRecursiveChildProcesses(3)
     //cproc := procfs.FindChildren( pid )
     //cproc,err  := procfs.ChildProcesses(pid)
     //fmt.Println("child array is: ",cproc)
     // fetch := func(proc Process) {
     //      if proc.Children == nil {
     //           return
     //      }

	// }
     // go fetch(v)
     for  i,e := range v.Children {
          fmt.Println("I cpid: ",i)
          fmt.Println("E Children pid", e.Pid)
          fmt.Println("E Command: ", e.Cmd)
		for  q,w := range e.Children {
			fmt.Println("W cpid: ",q)
			fmt.Println("W Children pid", w.Pid)
          	fmt.Println("W Command: ", w.Cmd)
		}

     }
//	 fmt.Println("Children", v.Children)

     //}
     fmt.Println("hello")
}
