include "Types.dfy"
include "ZooKeeper_ZKDatabase"


module ZooKeeper_Follower {
import opened ZooKeeper_Types
import opened ZooKeeper_ZKDatabase

datatype FollowerState = F_HANDSHAKE | F_SYNC | F_RUNNING

datatype Follower = Follower(
    my_id: nat,
    leader_id: nat,
    config: Config, // The cluster configuration. config[my_id] is my own endpoint, config[leader_id] is the leader endpoint
    zkdb: ZKDatabase,
    state: FollowerState 
)

predicate FollowerInit(s:Follower, my_id:nat, leader_id:nat, config:Config, zkdb: ZKDatabase)
{
    && s.my_id == my_id
    && s.leader_id == leader_id
    && s.config == config
    && s.zkdb == zkdb
    && s.state == F_HANDSHAKE
}

// predicate FollowerSend(s:Follower, s':Follower, )

}

