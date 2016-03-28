package env

type DevelopEnvironment struct {}

func (t *DevelopEnvironment) GetDBHost() string {
	return "localhost"
}

func (t *DevelopEnvironment) GetDBName() string {
	return "gosearch"
}
