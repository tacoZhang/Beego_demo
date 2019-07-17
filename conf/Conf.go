package conf

type Conf struct {
	MariaConf Maria
}

type Maria struct {
	Host     string
	Port     string
	UserName string
	PassWord string
	Db       string
}
