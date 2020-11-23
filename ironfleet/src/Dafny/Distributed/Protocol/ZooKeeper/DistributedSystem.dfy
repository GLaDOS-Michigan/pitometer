include "../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../Common/Framework/EnvironmentTCP.s.dfy"

include "Types.dfy"
include "ZKEnvironment.dfy"
include "ZKDatabase.dfy"
include "Follower.dfy"
include "Leader.dfy"
include "LearnerHandler.dfy"


module ZooKeeper_DistributedSystem {
import opened Common__SeqIsUniqueDef_i

import opened ZooKeeper_Types
import opened ZooKeeper_Environment
import opened EnvironmentTCP_s
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Follower
import opened ZooKeeper_Leader
import opened ZooKeeper_LearnerHandler

datatype QuorumPeer = LeaderPeer(leader:Leader) | FollowerPeer(follower:Follower)

datatype LS_State = LS_State(
    environment: ZKEnvironment,
    initialZkdbState: seq<ZKDatabase>,
    servers: map<EndPoint, QuorumPeer>
)


predicate LS_Init(s:LS_State, config:Config, f: int) {
    && f >= 1
    && LEnvironment_Init(s.environment, config)
    && |config| == |s.initialZkdbState| == 2*f + 1 // we will assign each server in config the corresponding db
    && SeqIsUnique(config)
    && InitialZkdbState_EmptyDiff(s.initialZkdbState)
    && (forall e :: e in config <==> e in s.servers)
        // this is the leader
    && s.servers[config[0]].LeaderPeer?
    && LeaderInit(s.servers[config[0]].leader, 0, config, s.initialZkdbState[0])
        // remaining nodes are followers
    && (forall i | 1 <= i < |config| :: 
            && s.servers[config[i]].FollowerPeer?
            && FollowerInit(s.servers[config[i]].follower, i, 0, config, s.initialZkdbState[i])
    )
}


predicate LS_NextOneServer(s:LS_State, s':LS_State, id:EndPoint, ios:seq<ZKIo>)
        requires id in s.servers;
{
    && id in s'.servers
    && s'.servers == s.servers[id := s'.servers[id]]
    && match s.servers[id] 
        case LeaderPeer(leader) => 
            && s'.servers[id].LeaderPeer?
            && LeaderNext(leader, s'.servers[id].leader, ios)
        case FollowerPeer(follower) => 
            && s'.servers[id].FollowerPeer?
            && FollowerNext(follower, s'.servers[id].follower, ios)
}


predicate LS_Next(s:LS_State, s':LS_State){
        LEnvironment_Next(s.environment, s'.environment)
    && if s.environment.nextStep.LEnvStepHostIos? && s.environment.nextStep.actor in s.servers then
            LS_NextOneServer(s, s', s.environment.nextStep.actor, s.environment.nextStep.ios)
    else
            s'.servers == s.servers
}



/*****************************************************************************************
*                             Assorted Initial zkdb States                               *
*****************************************************************************************/

/* Specifies a valid initial, on-disk copy of a zkdb that Zookeeper servers load into mem
* We first specify the situation where we should send empty diff */
predicate InitialZkdbState_EmptyDiff(zkdbs: seq<ZKDatabase>) {
    && (forall i | 0 <= i < |zkdbs| :: (
            && ZKDatabaseInit(zkdbs[i])
            && zkdbs[i].minCommittedLog == zkdbs[i].maxCommittedLog == NullZxid  // empty in-mem segment
        )
    ) && (forall i, j | 0 <= i < |zkdbs| && 0 <= j < |zkdbs| :: (
            zkdbs[i].commitLog == zkdbs[j].commitLog    // commit logs are identical
        )
    )
}

}