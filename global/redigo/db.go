package redigo

type dbRds struct {
}

func (d *dbRds) SelectDb(db int64) *Reply {
	c := (*Pool).Get()
	defer c.Close()
	return getReply(c.Do("select", db))
}
