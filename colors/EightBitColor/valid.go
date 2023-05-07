package twobitcolor

import (
  . "github.com/pandasoli/colorstrings/utils"
)


func valid_foreground(v uint8) bool { return UIntIn(uint(v), []uint { 38, 58 }) }
func valid_background(v uint8) bool { return UIntIn(uint(v), []uint { 48 }) }
