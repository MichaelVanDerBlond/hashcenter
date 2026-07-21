package workflow

func NewDefaultPipeline() *Engine {
	engine := NewEngine()

	engine.Add(NewAnalyzeStage())
	engine.Add(NewConvertStage())
	engine.Add(NewAttackStage())

	return engine
}
