include "Types.dfy"
include "ZKDatabase.dfy"


module ZooKeeper_Follower {
import opened ZooKeeper_Types
import opened ZooKeeper_ZKDatabase


// config[my_id] is my own endpoint, config[leader_id] is the leader endpoint
datatype Follower = Follower(
    my_id: nat,
    leader_id: nat,
    config: Config,
    zkdb: ZKDatabase,
    state: FollowerState 
)

datatype FollowerState = F_HANDSHAKE_A | F_HANDSHAKE_B | F_SYNC | F_RUNNING

predicate FollowerInit(s:Follower, my_id:nat, leader_id:nat, config:Config, zkdb: ZKDatabase)
{
    && s.my_id == my_id
    && s.leader_id == leader_id
    && s.config == config
    && s.zkdb == zkdb
    && s.state == F_HANDSHAKE_A
}

predicate FollowerNext(s:Follower, s':Follower, ios:seq<ZKIo>) {
    match s.state 
        case F_HANDSHAKE_A => SendMyInfo(s, s', ios)
        case F_HANDSHAKE_B => false // TODO
        case F_SYNC => false  // TODO
        case F_RUNNING => false  // TODO
}


/* State transition from F_HANDSHAKE_A -> F_HANDSHAKE_B */
predicate SendMyInfo(s:Follower, s':Follower, ios:seq<ZKIo>) {
    && s' == s.(state := F_HANDSHAKE_B)
    && |ios| == 1
    && 0 <= s.leader_id < |s.config|
    && ios[0].LIoOpSend?
    && var outbound_packet := ios[0].s;
        && outbound_packet.dst == s.config[s.leader_id]
        && outbound_packet.msg.FollowerInfo?
        && outbound_packet.msg.sid == s.my_id
        && if |s.zkdb.commitLog| == 0 
            then outbound_packet.msg.latestZxid == NullZxid
            else outbound_packet.msg.latestZxid == s.zkdb.commitLog[|s.zkdb.commitLog|-1]
}
}