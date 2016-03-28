package env

type TestEnvironment struct {}

func (t *TestEnvironment) GetDBHost() string {
	return "localhost"
}

func (t *TestEnvironment) GetDBName() string {
	return "test"
}
