package strmech

import "sync"

// Do NOT access these maps without first getting
// the lock on 'lockIntegerSeparatorTypeCode'.

var mIntegerSeparatorTypeCodeToString = map[IntegerSeparatorType]string{
	IntegerSeparatorType(0): "None",
	IntegerSeparatorType(1): "Thousands",
	IntegerSeparatorType(2): "IndiaNumbering",
	IntegerSeparatorType(3): "ChineseNumbering",
}

var mIntegerSeparatorTypeStringToCode = map[string]IntegerSeparatorType{
	"None":             IntegerSeparatorType(0),
	"Thousands":        IntegerSeparatorType(1),
	"IndiaNumbering":   IntegerSeparatorType(2),
	"ChineseNumbering": IntegerSeparatorType(3),
}

type IntegerSeparatorType int

var lockIntegerSeparatorTypeCode sync.Mutex
