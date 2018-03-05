package main

import (
	"testing"
)

//Test Read white list form config file
//测试读取白名单的函数
func TestReadConfig(t *testing.T) {
	var expectantResult = []string{"alpine:3.7"}
	result, err := ReadConfig("whiteList_test")
	if err != nil {
		t.Errorf("Read file error", err)
	}
	if stringSliceEqualBCE(expectantResult, result) {
		t.Log("Test Pass")
	} else {
		t.Errorf("Read file error,result is not equal to expectantResult")
	}
}

func TestInWhiteList(t *testing.T) {
	whiteList, _ = ReadConfig("whiteList_test")

	image1 := Image{
		RepoTags: "alpine:3.7",
	}
	if !inWhiteList(image1) {
		t.Errorf("Test Failed,InWhiteList is error")
	}

	image2 := Image{
		RepoTags: "alpine:3.6",
	}
	if inWhiteList(image2) {
		t.Errorf("Test Failed,InWhiteList is error")
	}

}

func TestStringSliceEqualBCE(t *testing.T) {
	var slideA = []string{"a", "b", "c"}
	var slideB = []string{"a", "b", "c"}

	if stringSliceEqualBCE(slideA, slideB) {
		t.Log("Test Pass")
		return
	}
	t.Errorf("Test Failed")
}

func stringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
