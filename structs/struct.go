package structs

type Workflow struct {
	Name string      `yaml:"name"`
	On   interface{} `yaml:"on"`
	Jobs Jobs        `yaml:"jobs"`
}

type Jobs struct {
	Deploy Deploy `yaml:"deploy"`
}

type Deploy struct {
	RunsOn string `yaml:"runs-on"`
	Steps  []Step `yaml:"steps"`
}

type Step struct {
	Uses string `yaml:"uses,omitempty"`
	Name string `yaml:"name,omitempty"`
	Env  Env    `yaml:"env,omitempty"`
	Run  string `yaml:"run,omitempty"`
}

type Env map[string]interface{}
