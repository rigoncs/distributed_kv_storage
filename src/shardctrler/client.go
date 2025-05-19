package shardctrler

import (
	"course/labrpc"
	"crypto/rand"
	"math/big"
)

//
// Shardctrler clerk
//

type Clerk struct {
	servers []*labrpc.ClientEnd
	// Your data here.
}

func nrand() int64 {
	max := big.NewInt(int64(1) << 62)
	bigx, _ := rand.Int(rand.Reader, max)
	x := bigx.Int64()
	return x
}

func MakeClerk(servers []*labrpc.ClientEnd) *Clerk {
	ck := new(Clerk)
	ck.servers = servers
	// Your code here.
	return nil
}

func (ck *Clerk) Query(num int) Config {
	// Your code here.
	return Config{}
}

func (ck *Clerk) Join(servers map[int][]string) {
	// Your code here.
}

func (ck *Clerk) Leave(gids []int) {
	// Your code here.
}

func (ck *Clerk) Move(shard, gid int) {
	// Your code here.
}
