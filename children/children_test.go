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

package children

import (
	"github.com/chair300/procfs/util"
	"testing"
)



func TestParsingChildren(t *testing.T) {
	// set the GLOBAL_SYSTEM_START
	util.GLOBAL_SYSTEM_START = 1388417200
	s, err := New("./testfiles", 38981)

	if err != nil {
		t.Fatal("Got error", err)
	}

	if s == nil {
		t.Fatal("children is missing")
	}

	// if s.Starttime.seconds() != 1388604586 {
		// t.Fatal("Start time is wrong, expected: 1388604586", s.Starttime.EpochSeconds)
	// }
	t.Log(s)
	var tobe bool = false
	for _, b := range s {
		if b == 39893 {
			tobe = true;
			t.Log("Contains 39893")
		}
	}
	if !tobe {
		t.Fatal("Does not contain the child proc id")
	}

}

func BenchmarkParsingChildren(b *testing.B) {
	// set the GLOBAL_SYSTEM_START
	util.GLOBAL_SYSTEM_START = 1388417200
	s, err := New("./testfiles", 38981)

	if err != nil {
		b.Fatal("Got error", err)
	}

	if s == nil {
		b.Fatal("children is missing")
	}

	// if s.Starttime.seconds() != 1388604586 {
		// t.Fatal("Start time is wrong, expected: 1388604586", s.Starttime.EpochSeconds)
	// }

	var tobe bool = false
	for _, b := range s {
		if b == 39893 {
			tobe = true;
			//b.Log("Contains 39893")
		}
	}
	if !tobe {
		b.Fatal("Does not contain the child proc id")
	}

}
