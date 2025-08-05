package Collection

type Collection interface {
	Insert(int)
	ForEach(func(int))
}
