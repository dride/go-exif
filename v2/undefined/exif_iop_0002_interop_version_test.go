package exifundefined

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/dsoprea/go-logging"

	"github.com/dride/go-exif/v2/common"
)

func TestTag0002InteropVersion_String(t *testing.T) {
	ut := Tag0002InteropVersion{"abc"}
	s := ut.String()

	if s != "abc" {
		t.Fatalf("String not correct: [%s]", s)
	}
}

func TestCodec0002InteropVersion_Encode(t *testing.T) {
	s := "abc"
	ut := Tag0002InteropVersion{s}

	codec := Codec0002InteropVersion{}

	encoded, unitCount, err := codec.Encode(ut, exifcommon.TestDefaultByteOrder)
	log.PanicIf(err)

	if bytes.Equal(encoded, []byte(s)) != true {
		t.Fatalf("Encoded bytes not correct: %v", encoded)
	} else if unitCount != uint32(len(s)) {
		t.Fatalf("Unit-count not correct: (%d)", unitCount)
	}
}

func TestCodec0002InteropVersion_Decode(t *testing.T) {
	s := "abc"
	ut := Tag0002InteropVersion{s}

	encoded := []byte(s)

	rawValueOffset := encoded

	valueContext := exifcommon.NewValueContext(
		"",
		0,
		uint32(len(encoded)),
		0,
		rawValueOffset,
		nil,
		exifcommon.TypeUndefined,
		exifcommon.TestDefaultByteOrder)

	codec := Codec0002InteropVersion{}

	value, err := codec.Decode(valueContext)
	log.PanicIf(err)

	if reflect.DeepEqual(value, ut) != true {
		t.Fatalf("Decoded value not correct: %s\n", value)
	}
}
