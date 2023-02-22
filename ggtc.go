package ggtc

func Doot() string {
	// shell function to test module existence
	return "snoot: dooted"
}

type SimulationResult interface {
	GetInput() Input
	GetFitness() int
	GetMetadata() interface{}
}

type Input interface{}

type SimulationConfig map[string]interface{}

type Simulation interface {
	Process(in Input, sc SimulationConfig) SimulationResult
}

type EngineConfig struct {
	SimulationConfig SimulationConfig
	Seed             string
	Generations      int
}

type EngineResult struct {
	FittestResults []SimulationResult
}

type Engine interface {
	Configure(ec EngineConfig)
	Run() EngineResult
}
