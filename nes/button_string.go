// generated by stringer -type=Button; DO NOT EDIT

package nes

import "fmt"

const _Button_name = "ABSelectStartUpDownLeftRightOne"

var _Button_index = [...]uint8{0, 1, 2, 8, 13, 15, 19, 23, 28, 31}

func (i Button) String() string {
	if i+1 >= Button(len(_Button_index)) {
		return fmt.Sprintf("Button(%d)", i)
	}
	return _Button_name[_Button_index[i]:_Button_index[i+1]]
}
