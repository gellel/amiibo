package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engAmiiboFileName = "eng-amiibo-test.json"

func testENGAmiibo(t *testing.T, v *amiibo.ENGAmiibo) {
	if name := v.GetNameAlternative(); len(name) == 0 {
		t.Fatal("len((ENGAmiib).GetNamespace()) == 0")
	}
}

func TestENGAmiibo(t *testing.T) {
	var err error
	var ENGAmiibo = amiibo.ENGAmiibo{}
	_, err = amiibo.WriteENGAmiibo(filefolder, engAmiiboFileName, &ENGAmiibo)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGAmiibo(filefolder, engAmiiboFileName)
	if err != nil {
		t.Fatal(err)
	}
}
