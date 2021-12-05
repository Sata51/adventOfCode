package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

// Data holder struct
type Data struct {
	content []byte
}

//Load file
func Load(year int, filename string) Data {
	dat, err := ioutil.ReadFile(fmt.Sprintf("../../input/%d/%s", year, filename))
	panicError(err)
	return Data{content: dat}
}

//String retruns string representation
func (d Data) String() string {
	return string(d.content)
}

//ToIntSlice returns a slice of int from Data
func (d Data) ToIntSlice() []int {
	retVal := make([]int, 0)
	for _, v := range strings.Split(d.String(), "\n") {
		retVal = append(retVal, MustParseInt(v))
	}
	return retVal
}

// ToStringSlice returns a slice of string from Data
func (d Data) ToStringSlice() []string {
	retVal := make([]string, 0)
	retVal = append(retVal, strings.Split(d.String(), "\n")...)
	return retVal
}
