package doctor

type Check interface {
	Name() string
	Run() Result
}
