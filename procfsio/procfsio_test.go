package procfsio

import (
	"github.com/chair300/procfs/util"
//	"strconv"
	"testing"
//	"log"
)


func TestParsingIo(t *testing.T) {
	// set the GLOBAL_SYSTEM_START
	util.GLOBAL_SYSTEM_START = 1388417200
	s, err := New("./testfiles/io")

	if err != nil {
		t.Fatal("Got error", err)
	}

	if s == nil {
		t.Fatal("stat is missing")
	}
	//log.Println("io is : ", s)
	if s.rchar != 119010 {
		t.Fatal("read char should be: 119010, instead",  s.rchar)
	}

}
