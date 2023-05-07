package fourbitcolor


func valid_foreground(v uint8) bool { return 90 <= v && v <= 97 }
func valid_background(v uint8) bool { return 100 <= v && v <= 107 }
