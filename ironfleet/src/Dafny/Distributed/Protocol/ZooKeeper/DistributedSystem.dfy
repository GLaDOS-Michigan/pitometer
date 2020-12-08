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


predicate LS_Init(config:Config, s:LS_State, f: int) {
    && f >= 1
    && LEnvironment_Init(config, s.environment)
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


predicate LS_NextOneServer(s:LS_State, s':LS_State, actor:EndPoint, ios:seq<ZKIo>)
    requires actor in s.servers
{
    && s.environment.nextStep == LEnvStepHostIos(actor, ios)
    && actor in s'.servers
    && s'.servers == s.servers[actor := s'.servers[actor]]
    && match s.servers[actor] 
        case LeaderPeer(leader) => 
            && s'.servers[actor].LeaderPeer?
            && LeaderNext(leader, s'.servers[actor].leader, ios)
        case FollowerPeer(follower) => 
            && s'.servers[actor].FollowerPeer?
            && FollowerNext(follower, s'.servers[actor].follower, ios)
}


predicate LS_Next(s:LS_State, s':LS_State){
    && LEnvironment_Next(s.environment, s'.environment)
    && (exists ep, ios :: ep in s.servers && LS_NextOneServer(s, s', ep, ios))
    && s'.initialZkdbState == s.initialZkdbState
}
}
