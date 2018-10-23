//
// children.children describes data in /proc/<pid>/tasks/<pid>/children.
//
// Use statm.New() to create a new stat.Statm object
// from data in a path.
//
package children

//
// Copyright Christopher Harrison (harrison@glsan.com)
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//

import (
	"io/ioutil"
	"strings"
	"strconv"
	"path"
	"log"

)


func New(prefix string, spid int64) ([]int64, error) {
	pids := *int64
	files, err := ioutil.ReadFile(path.Join(prefix, strconv.Itoa(spid), "task", strconv.Itoa(spid), "children"))
	if err != nil {
		return nil, err
	}
	mpids :=  strings.Fields(string(files))

	//todo := len(pids)

	// create a goroutine that asynchronously processes all /proc/<pid> entries
	for _, str := range mpids {
		//pid, err := strconv.Atoi(str)
		pid, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			// TODO: bubble errors up if requested
			log.Println("Failure loading process pid: ", pid, err)
			//done <- nil
		}else{
			pids[pid] = pid
		}
	}
	return pids, err
}
