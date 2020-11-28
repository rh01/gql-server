package store

import mgo "gopkg.in/mgo.v2"

// Database 存储数据库信息
type Database struct {
	Session *mgo.Session
	Name    string
}

// NewDatabase Create a new base datasource
func NewDatabase(dbname string, session *mgo.Session) Database {
	return Database{Name: dbname, Session: session}
}

// Report 存储多个结构体
type Report struct {
	Vcs     Vcs
	Sysconf Sysconf
}

// InitReportStore ：初始化Report项目的存储
func (db Database) InitReportStore() Report {
	// create...
	return Report{}
}
