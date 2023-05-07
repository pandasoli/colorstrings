package field


type Field struct {
  value uint8
  props []uint8
}

func (self Field) GetValue() uint8 { return self.value }
func (self Field) GetProps() []uint8 { return self.props }

func NewField(value uint8, props []uint8) Field {
  return Field { value, props }
}
