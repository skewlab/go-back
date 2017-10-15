/*
*/

package util

import (
	"fmt"
	"io/ioutil"
)

func ReadFile( file string ) string {
	text, err := ioutil.ReadFile( file )
	if err != nil {
		fmt.Print( err )
	}
	return string( text )
}
