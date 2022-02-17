package redigo

import "bankroll/global"

type dtype struct {
	String stringRds
	List   listRds
	Hash   hashRds
	Key    keyRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
	Db     dbRds
}
var Dtype = new(dtype)

var Pool = &global.Redigo


