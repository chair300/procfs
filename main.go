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

package main

import (
     "fmt"
	"github.com/chair300/procfs"

)


func main() {

     // fetch all processes from /proc
     // returns a map of pid -> *Process
     procs, err := procfs.Processes(false);
     //fmt.Println(procs)
     if err != nil {}
     //fmt.Println("processes", processes)
     for k, v := range procs {
          fmt.Println("pid:", k)
          fmt.Println("Command:", v.Cmd)
          s, err := v.Statm()
          if err == nil {
               fmt.Println("size:", s.Size)
          }
          //fmt.Println("size:", s.Size)
	  for elem := range v.Children {
		fmt.Println("Children", elem.Pid)
    	  }
	 fmt.Println("Children", v.Children)

     }
     fmt.Println("hello")
}
