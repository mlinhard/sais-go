package sais

// #cgo CFLAGS: -g -Wall
// #include "sais.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func sais(c_T *C.uchar, c_SA *C.int, c_n C.int) error {
	r := C.sais(c_T, c_SA, c_n)
	if r != 0 {
		return fmt.Errorf("sais function returned %v", r)
	}
	return nil
}

func sais_int(c_T *C.int, c_SA *C.int, c_n C.int, c_k C.int) error {
	r := C.sais_int(c_T, c_SA, c_n, c_k)
	if r != 0 {
		return fmt.Errorf("sais_int function returned %v", r)
	}
	return nil
}

func TextSuffixArray(text string) ([]int, error) {
	SA := make([]int, len(text))
	SA32 := make([]int32, len(text))
	textBytes := []byte(text)
	err := Sais32(textBytes, SA32)
	if err != nil {
		return nil, err
	}
	for i, val := range SA32 {
		SA[i] = (int)(val)
	}
	return SA, nil
}

func SuffixArray(T []byte) ([]int, error) {
	SA := make([]int, len(T))
	SA32 := make([]int32, len(T))
	err := Sais32(T, SA32)
	if err != nil {
		return nil, err
	}
	for i, val := range SA32 {
		SA[i] = (int)(val)
	}
	return SA, nil
}

/* find the suffix array SA of T[0..n-1]
   use a working space (excluding T and SA) of at most 2n+O(lg n) */
func Sais32(T []byte, SA []int32) error {
	return sais(
		(*C.uchar)(unsafe.Pointer(&T[0])),
		(*C.int)(unsafe.Pointer(&SA[0])),
		C.int(len(SA)))
}

/* find the suffix array SA of T[0..n-1]
   use a working space (excluding T and SA) of at most 2n+O(lg n) */
func Sais64(T []byte, SA []int64) error {
	return sais(
		(*C.uchar)(unsafe.Pointer(&T[0])),
		(*C.int)(unsafe.Pointer(&SA[0])),
		C.int(len(SA)))
}

/* find the suffix array SA of T[0..n-1] in {0..k-1}^n
use a working space (excluding T and SA) of at most MAX(4k,2n) */
func SaisInt32(T []int32, SA []int32, k int32) error {
	if len(T) != len(SA) {
		return fmt.Errorf("Length of T and SA arrays must be equal")
	}
	return sais_int(
		(*C.int)(unsafe.Pointer(&T[0])),
		(*C.int)(unsafe.Pointer(&SA[0])),
		C.int(len(SA)),
		C.int(k))
}

/* find the suffix array SA of T[0..n-1] in {0..k-1}^n
use a working space (excluding T and SA) of at most MAX(4k,2n) */
func SaisInt64(T []int64, SA []int64, k int64) error {
	if len(T) != len(SA) {
		return fmt.Errorf("Length of T and SA arrays must be equal")
	}
	return sais_int(
		(*C.int)(unsafe.Pointer(&T[0])),
		(*C.int)(unsafe.Pointer(&SA[0])),
		C.int(len(SA)),
		C.int(k))
}

/* burrows-wheeler transform */
/*  int
sais_bwt(const unsigned char *T, unsigned char *U, int *A, int n);
int
sais_int_bwt(const int *T, int *U, int *A, int n, int k);*/
