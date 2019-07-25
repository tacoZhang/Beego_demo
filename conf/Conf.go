package conf

// 时间格式化字符串
const SysTimeform string = "2006-01-02 15:04:05"
const SysTimeformShort string = "2006-01-02"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

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
