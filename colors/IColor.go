package colors


type IColor interface {
  GetKind() ColorKind
  Join() string
}
