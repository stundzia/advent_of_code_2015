package day5

import (
	"testing"
)

func TestIsNice(t *testing.T) {
	nice := isNice("ugknbfddgicrmopn")
	if !nice {
		t.Errorf("`ugknbfddgicrmopn` should be nice")
	}

	nice = isNice("aaa")
	if !nice {
		t.Errorf("`aaa` should be nice")
	}

	nice = isNice("jchzalrnumimnmhp")
	if nice {
		t.Errorf("`jchzalrnumimnmhp` should be naughty")
	}

	nice = isNice("haegwjzuvuyypxyu")
	if nice {
		t.Errorf("`haegwjzuvuyypxyu` should be naughty")
	}

	nice = isNice("dvszwmarrgswjxmb")
	if nice {
		t.Errorf("`dvszwmarrgswjxmb` should be naughty")
	}
}


func TestIsNice2(t *testing.T) {
	nice := isNice2("qjhvhtzxzqqjkmpb")
	if !nice {
		t.Errorf("`qjhvhtzxzqqjkmpb` should be nice")
	}

	nice = isNice2("haha")
	if !nice {
		t.Errorf("`haha` should be nice")
	}

	nice = isNice2("aabcdefegaa")
	if !nice {
		t.Errorf("`aabcdefegaa` should be nice")
	}

	nice = isNice2("xxyxx")
	if !nice {
		t.Errorf("`xxyxx` should be nice")
	}

	nice = isNice2("uurcxstgmygtbstg")
	if nice {
		t.Errorf("`uurcxstgmygtbstg` should be naughty")
	}

	nice = isNice2("ieodomkazucvgmuy")
	if nice {
		t.Errorf("`ieodomkazucvgmuy` should be naughty")
	}

	nice = isNice2("ierrodasaomkazucvgmuy")
	if nice {
		t.Errorf("`ieodomkazucvgmuy` should be naughty")
	}

	nice = isNice2("aaaa")
	if !nice {
		t.Errorf("`aaaa` should be nice")
	}

	nice = isNice2("ttt")
	if nice {
		t.Errorf("`ttt` should be naughty")
	}

	nice = isNice2("zztdcqzyqddaazdjp")
	if nice {
		t.Errorf("`zztdcqzyqddaazdjp` should be naughty")
	}

}