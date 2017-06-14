package main

import (
	"fmt"
	"testing"
)

func TestGetSetterGetter(t *testing.T) {
	testObjects := []struct {
		Input    string
		Prefix   string
		Expected string
	}{
		{
			Input:  "myVar",
			Prefix: "public",
		},

		{
			Input:  "room",
			Prefix: "private",
		},
	}

	for _, testObject := range testObjects {
		result := getSetterGetter(testObject.Prefix, testObject.Input)
		fmt.Println(result)
	}
}

func TestGetVariables(t *testing.T) {
	testObjects := []struct {
		Datas    []byte
		Expected []string
	}{
		{
			Datas: []byte(`<?php
				class SomeClass {
    				private $myVar;
    				private $newVaria;
    				private $newVar;

				}`),
			Expected: []string{"myVar", "newVaria", "newVar"},
		},

		{
			Datas: []byte(`<?php
				class SomeClass {
    				private $hello;
    				private $Wkwkw;
    				private $eheHaha;

				}`),
			Expected: []string{"hello", "Wkwkw", "eheHaha"},
		},
	}

	for _, testObject := range testObjects {
		actual := getVarNames(testObject.Datas)
		if !eq(actual, testObject.Expected) {
			t.Errorf("actual = %v, expected = %v\n", actual, testObject.Expected)
		}

	}
}

func eq(a, b []string) bool {
	if len(a) != len(b) {
		fmt.Println("masuk gk")
		return false
	}

	if a == nil || b == nil {
		return false
	}

	for index, value := range a {
		if value != b[index] {
			return false
		}
	}

	return true

}
