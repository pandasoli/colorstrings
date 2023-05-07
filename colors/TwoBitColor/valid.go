package twobitcolor


func valid_foreground(v uint8) bool { return 30 <= v && v <= 37 || v == 39 }
func valid_background(v uint8) bool { return 40 <= v && v <= 47 || v == 49 }
