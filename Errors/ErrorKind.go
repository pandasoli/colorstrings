package errors


type ErrorKind uint8

const (
  NO_ERROR ErrorKind = iota
  WRONG_COLOR  // 1, 37, 49, 100, 101...
  WRONG_PREFIX // The next field after 38 or 48
  WRONG_PROPS  // Props
  WRONG_FIELDS // Expected more fields
  NO_PREFIX    // Missing prefix
)
