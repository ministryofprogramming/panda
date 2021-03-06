package plugins

// Command that the plugin supplies
type Command struct {
	// Name "foo"
	Name string `json:"name"`
	// UseCommand "bar"
	UseCommand string `json:"use_command"`
	// PandaCommand "generate"
	PandaCommand string `json:"panda_command"`
	// Description "generates a foo"
	Description string   `json:"description"`
	Aliases     []string `json:"aliases"`
	Binary      string   `json:"-"`
}

// Commands is a slice of Command
type Commands []Command
