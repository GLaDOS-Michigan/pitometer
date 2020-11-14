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
    accepted_epoch: int,
    state: FollowerState 
)

datatype FollowerState = F_HANDSHAKE_A | F_HANDSHAKE_B | F_SYNC | F_RUNNING | F_ERROR

predicate FollowerInit(s:Follower, my_id:nat, leader_id:nat, config:Config, zkdb: ZKDatabase)
{
    && s.my_id == my_id
    && s.leader_id == leader_id
    && s.config == config
    && s.zkdb == zkdb
    && s.accepted_epoch == -1
    && s.state == F_HANDSHAKE_A
}

predicate FollowerNext(s:Follower, s':Follower, ios:seq<ZKIo>) {
    match s.state 
        case F_HANDSHAKE_A => SendMyInfo(s, s', ios)
        case F_HANDSHAKE_B => false // TODO
        case F_SYNC => false  // TODO
        case F_RUNNING => false  // TODO
        case F_ERROR => s' == s  // TODO
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


/* State transition from F_HANDSHAKE_B -> F_SYNC */
predicate AcceptNewEpoch(s:Follower, s':Follower, ios:seq<ZKIo>) {
    // state and epoch are the only properties that change in this transition
    && s' == s.(state := s'.state, accepted_epoch := s'.accepted_epoch)  
    && |ios| >= 1
    && 0 <= s.leader_id < |s.config|
    && ios[0].LIoOpReceive?
    && ios[0].r.msg.LeaderInfo?
    && ios[0].r.msg.sid == s.leader_id
    && (if ios[0].r.msg.newZxid.epoch <  s.accepted_epoch 
        then && s'.state == F_ERROR
             && |ios| == 1
        else (
            && s'.state == F_SYNC
            && s'.accepted_epoch == ios[0].r.msg.newZxid.epoch
            && |ios| == 2
            && ios[1].LIoOpSend?
            && ios[1].s.dst == s.config[s.leader_id]
            && ios[1].s.msg.AckEpoch?
            && ios[1].s.msg.sid == s.my_id
            && ios[1].s.msg.lastLoggedZxid == getLastLoggedZxid(s.zkdb)
            && if ios[0].r.msg.newZxid.epoch ==  s.accepted_epoch
                then ios[1].s.msg.lastAcceptedEpoch == -1
                else ios[1].s.msg.lastAcceptedEpoch == ios[0].r.msg.newZxid.epoch
        )
    )
}
}