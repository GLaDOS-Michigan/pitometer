include "../../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "../Timestamps/TimestampedType.dfy"
include "TimestampedLS.dfy"

include "../Types.dfy"
include "../DistributedSystem.dfy"
include "../ZKEnvironment.dfy"
include "../ZKDatabase.dfy"
include "../Follower.dfy"
include "../Leader.dfy"
include "../LearnerHandler.dfy"
include "Definitions.dfy"


/* This module contains some common lemmas */
module Zookeeper_Commons {
import opened Common__SeqIsUniqueDef_i
import opened ZKTimestamp
import opened ZooKeeper_Types
import opened ZooKeeper_Environment
import opened ZooKeeper_DistributedSystem
import opened EnvironmentTCP_s
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Follower
import opened ZooKeeper_Leader
import opened ZooKeeper_LearnerHandler
import opened Zookeeper_Performance_Definitions
import opened ZooKeeper_TimestampedDistributedSystem


predicate SentPacketsSet_Property(tls:TLS_State, tls':TLS_State, id:EndPoint, tios:seq<TZKIo>)
    requires id in tls.t_servers
    requires TLS_Next(tls, tls')
    requires TLS_NextOneServer(tls, tls', id, tios)
{
    tls'.t_environment.sentPackets == 
    tls.t_environment.sentPackets + (set tio : TimestampedLIoOp | tio in tios && tio.LIoOpSend? :: tio.s)
}


lemma lemma_SentPacketsSet_Property(tls:TLS_State, tls':TLS_State, id:EndPoint, tios:seq<TZKIo>)
    requires id in tls.t_servers
    requires TLS_Next(tls, tls')
    requires TLS_Next(tls, tls')
    requires TLS_NextOneServer(tls, tls', id, tios)
    ensures SentPacketsSet_Property(tls, tls', id, tios)
{}


lemma lemma_Math_Inequality(a:nat, b:nat)
    requires a <= b + 1
    ensures b-a+1 >= 0
{}


}
