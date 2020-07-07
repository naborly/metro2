// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzzreader

import (
	"reflect"

	"github.com/moov-io/metro2/pkg/file"
)

// Return codes (from go-fuzz docs)
//
// The function must return 1 if the fuzzer should increase priority
// of the given input during subsequent fuzzing (for example, the input is
// lexically correct and was parsed successfully); -1 if the input must not be
// added to corpus even if gives new coverage; and 0 otherwise; other values are
// reserved for future use.
func Fuzz(data []byte) int {
	f, err := file.NewFile(string(data))
	if err != nil {
		return -1
	}

	if err := f.Validate(); err != nil {
		return 0
	}

	gendata := []byte(f.String())
	if !reflect.DeepEqual(data, gendata) {
		return 0
	}

	// If we're missing a record the file is close, but we should continue around
	// that input value.
	if record, _ := f.GetRecord(file.HeaderRecordName); record == nil {
		return 0
	}
	if record, _ := f.GetRecord(file.TrailerRecordName); record == nil {
		return 0
	}
	if records := f.GetDataRecords(); len(records) == 0 {
		return 0
	}

	// Prioritize generated files with header, trailer, and data records.
	return 1
}