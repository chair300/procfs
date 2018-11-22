package procfsio

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

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

//
// io is a parser for /proc/<pid>/id.
//
//
type Procfsio struct {
	rchar      int64
	wchar       int64
	syscr       int64
	syscw        int64
	read_bytes    int64
	write_bytes        int64
	cancelled_write_bytes      int64
}


func New(path string) (*Procfsio, error) {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(buf), "\n")
	info := &Procfsio{}

	for i := 0; i < len(lines); i++ {

		lines[i] = strings.TrimSpace(lines[i])
 		fields := strings.Fields(lines[i])
		if strings.HasPrefix(lines[i], "rchar:") {
		//if i == 0 {
			if info.rchar,err = strconv.ParseInt(fields[1],10,64); err != nil {
	 			return nil, fmt.Errorf("Unable to parse rchar %s: %v", fields[1], err)
	 		}
		}

		if strings.HasPrefix(lines[i], "wchar:") {
		//if i == 1 {
			if info.wchar,err = strconv.ParseInt(fields[1],10,64); err != nil {
				return nil, fmt.Errorf("Unable to parse wchar %s: %v", fields[1], err)
			}
		}

		if strings.HasPrefix(lines[i], "syscr:") {
		//if i == 2 {
			if info.syscr,err = strconv.ParseInt(fields[1],10,64); err != nil {
				return nil, fmt.Errorf("Unable to parse syscr %s: %v", fields[1], err)
			}
		}

		if strings.HasPrefix(lines[i], "syscw:") {
		//if i == 3 {
			if info.syscw,err = strconv.ParseInt(fields[1],10,64); err != nil {
				return nil, fmt.Errorf("Unable to parse syscw %s: %v", fields[1], err)
			}
		}

		if strings.HasPrefix(lines[i], "read_bytes:") {
		//if i == 4 {
			if info.read_bytes,err = strconv.ParseInt(fields[1],10,64); err != nil {
				return nil, fmt.Errorf("Unable to parse read_bytes %s: %v", fields[1], err)
			}
		}

		if strings.HasPrefix(lines[i], "write_bytes:") {
		//if i == 5 {
			if info.write_bytes,err = strconv.ParseInt(fields[1],10,64); err != nil {
				return nil, fmt.Errorf("Unable to parse write_bytes %s: %v", fields[1], err)
			}
		}

		if strings.HasPrefix(lines[i], "cancelled_write_bytes:") {
		//if i == 6 {
			if info.cancelled_write_bytes,err = strconv.ParseInt(fields[1],10,64); err != nil {
				return nil, fmt.Errorf("Unable to parse cancelled_write_bytes %s: %v", fields[1], err)
			}
		}
	}
	return info, err
}
