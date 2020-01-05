package cmd

type ServerCommand struct {
	Port string `long:"port" env:"BISCUIT_PORT" short:"p" default:"8080" description:"biscuit server port"`
}

func (s *ServerCommand) Execute() {

}
