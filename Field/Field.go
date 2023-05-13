package field


type Field struct {
  Value uint8
  Props []uint8
}

func NewField(value uint8, props []uint8) Field {
  return Field { value, props }
}
