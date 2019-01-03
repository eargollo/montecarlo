package report

// New initializes a simulation object
func New(name string, api string, token string) Simulation {
	return Simulation{name: name, api: api, token: token}
}

// Simulation represents a MonteCarlo simulation
type Simulation struct {
	name  string
	api   string
	token string
}
