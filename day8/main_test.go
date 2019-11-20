package day8

import "testing"

func TestGetStringCodeCharCount(t *testing.T) {
	count := GetStringCodeCharCount(`"abc"`)
	if count != 5 {
		t.Fatalf("Expected count to be 5, got %d", count)
	}
	count = GetStringCodeCharCount(`"a\"bwe\"awe\"ffc"`)
	if count != 18 {
		t.Fatalf("Expected count to be 18, got %d", count)
	}
	count = GetStringCodeCharCount(`"\x27"`)
	if count != 6 {
		t.Fatalf("Expected count to be 6, got %d", count)
	}
	count = GetStringCodeCharCount(`"aaa\"aaa"`)
	if count != 10 {
		t.Fatalf("Expected count to be 10, got %d", count)
	}
	count = GetStringCodeCharCount(`"aaa\"\\"aaa"`)
	if count != 13 {
		t.Fatalf("Expected count to be 13, got %d", count)
	}
	count = GetStringCodeCharCount(`"vam\"yr\"q\x92o\x4ebeqe\\"`)
	if count != 27 {
		t.Fatalf("Expected count to be 27, got %d", count)
	}
}

func TestGetStringMemoryCharacterCount(t *testing.T) {
	count := GetStringMemoryCharacterCount(`"abc"`)
	if count != 3 {
		t.Fatalf("Expected count to be 3, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"a\"bwe\"awe\"ffc"`)
	if count != 13 {
		t.Fatalf("Expected count to be 13, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"\x27"`)
	if count != 1 {
		t.Fatalf(
			"Expected count to be 1, got %d",
			count,
			)
	}
	count = GetStringMemoryCharacterCount(`"\x27\xbe\\"`)
	if count != 3 {
		t.Fatalf(
			"Expected count to be 3, got %d",
			count,
		)
	}
	count = GetStringMemoryCharacterCount(`"aaa\"aaa"`)
	if count != 7 {
		t.Fatalf("Expected count to be 7, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"aaa\"\\"aaa"`)
	if count != 9 {
		t.Fatalf("Expected count to be 9, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"bvl\\\"noseec"`)
	if count != 11 {
		t.Fatalf("Expected count to be 11, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"vam\"yr\"q\x92o\x4ebeqe\\"`)
	if count != 16 {
		t.Fatalf("Expected count to be 16, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"\\lcomf\\s"`)
	if count != 8 {
		t.Fatalf("Expected count to be 8, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"g\x4ffxhq\\\xc1rw"`)
	if count != 10 {
		t.Fatalf("Expected count to be 10, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"\"\\fdktlp"`)
	if count != 8 {
		t.Fatalf("Expected count to be 8, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"vm\"yr\"q\x92o\x4eqe\\"`)
	if count != 13 {
		t.Fatalf("Expected count to be 13, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"\\xgicuuh"`)
	if count != 8 {
		t.Fatalf("Expected count to be 8, got %d", count)
	}
	count = GetStringMemoryCharacterCount(`"ly\\m\"pvnm\\xmyl"`)
	if count != 14 {
		t.Fatalf("Expected count to be 14, got %d", count)
	}
}

func TestGetEncodedLen(t *testing.T) {
	count := GetEncodedLen(`"abc"`)
	if count != 9 {
		t.Fatalf("Expected count to be 9, got %d", count)
	}
	count = GetEncodedLen(`""`)
	if count != 6 {
		t.Fatalf("Expected count to be 6, got %d", count)
	}
	count = GetEncodedLen(`"aaa\"aaa"`)
	if count != 16 {
		t.Fatalf("Expected count to be 16, got %d", count)
	}
	count = GetEncodedLen(`"\x27"`)
	if count != 11 {
		t.Fatalf("Expected count to be 11, got %d", count)
	}
	count = GetEncodedLen(`"aaa\"a\xebaa"`)
	if count != 21 {
		t.Fatalf("Expected count to be 21, got %d", count)
	}
}