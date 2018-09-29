// Copyright 2018 The MATRIX Authors as well as Copyright 2014-2017 The go-ethereum Authors
// This file is consisted of the MATRIX library and part of the go-ethereum library.
//
// The MATRIX-ethereum library is free software: you can redistribute it and/or modify it under the terms of the MIT License.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, 
//and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject tothe following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, 
//WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISINGFROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
//OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
// generated by stringer -output vt_string.go -type VT; DO NOT EDIT

package ole

import "fmt"

const (
	_VT_name_0 = "VT_EMPTYVT_NULLVT_I2VT_I4VT_R4VT_R8VT_CYVT_DATEVT_BSTRVT_DISPATCHVT_ERRORVT_BOOLVT_VARIANTVT_UNKNOWNVT_DECIMAL"
	_VT_name_1 = "VT_I1VT_UI1VT_UI2VT_UI4VT_I8VT_UI8VT_INTVT_UINTVT_VOIDVT_HRESULTVT_PTRVT_SAFEARRAYVT_CARRAYVT_USERDEFINEDVT_LPSTRVT_LPWSTR"
	_VT_name_2 = "VT_RECORDVT_INT_PTRVT_UINT_PTR"
	_VT_name_3 = "VT_FILETIMEVT_BLOBVT_STREAMVT_STORAGEVT_STREAMED_OBJECTVT_STORED_OBJECTVT_BLOB_OBJECTVT_CFVT_CLSID"
	_VT_name_4 = "VT_BSTR_BLOBVT_VECTOR"
	_VT_name_5 = "VT_ARRAY"
	_VT_name_6 = "VT_BYREF"
	_VT_name_7 = "VT_RESERVED"
	_VT_name_8 = "VT_ILLEGAL"
)

var (
	_VT_index_0 = [...]uint8{0, 8, 15, 20, 25, 30, 35, 40, 47, 54, 65, 73, 80, 90, 100, 110}
	_VT_index_1 = [...]uint8{0, 5, 11, 17, 23, 28, 34, 40, 47, 54, 64, 70, 82, 91, 105, 113, 122}
	_VT_index_2 = [...]uint8{0, 9, 19, 30}
	_VT_index_3 = [...]uint8{0, 11, 18, 27, 37, 55, 71, 85, 90, 98}
	_VT_index_4 = [...]uint8{0, 12, 21}
	_VT_index_5 = [...]uint8{0, 8}
	_VT_index_6 = [...]uint8{0, 8}
	_VT_index_7 = [...]uint8{0, 11}
	_VT_index_8 = [...]uint8{0, 10}
)

func (i VT) String() string {
	switch {
	case 0 <= i && i <= 14:
		return _VT_name_0[_VT_index_0[i]:_VT_index_0[i+1]]
	case 16 <= i && i <= 31:
		i -= 16
		return _VT_name_1[_VT_index_1[i]:_VT_index_1[i+1]]
	case 36 <= i && i <= 38:
		i -= 36
		return _VT_name_2[_VT_index_2[i]:_VT_index_2[i+1]]
	case 64 <= i && i <= 72:
		i -= 64
		return _VT_name_3[_VT_index_3[i]:_VT_index_3[i+1]]
	case 4095 <= i && i <= 4096:
		i -= 4095
		return _VT_name_4[_VT_index_4[i]:_VT_index_4[i+1]]
	case i == 8192:
		return _VT_name_5
	case i == 16384:
		return _VT_name_6
	case i == 32768:
		return _VT_name_7
	case i == 65535:
		return _VT_name_8
	default:
		return fmt.Sprintf("VT(%d)", i)
	}
}