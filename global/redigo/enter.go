package redigo_pack

import "bankroll/global"

type redigoPack struct {
	String stringRds
	List   listRds
	Hash   hashRds
	Key    keyRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
	Db     dbRds
}
var RedigoConn = new(redigoPack)

var Pool = global.GVA_REDIS

