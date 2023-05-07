package utils


func IntIn(val int, list []int) bool {
  for _, item := range list {
    if item == val { return true }
  }

  return false
}

func UIntIn(val uint, list []uint) bool {
  for _, item := range list {
    if item == val { return true }
  }

  return false
}
