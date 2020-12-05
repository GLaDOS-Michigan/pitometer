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


/* This module contains invariants that have nothing to do with performance */
module Zookeeper_EmptyDiffInvariants {
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


/*****************************************************************************************
*                                    Empty Diff facts                                    *
*****************************************************************************************/

predicate EmptyDiff_Invariant(tls:TLS_State) {
    forall pkt | pkt in tls.t_environment.sentPackets ::
        && !pkt.msg.v.SyncSNAP?
        && !pkt.msg.v.SyncTRUNC?
        && !pkt.msg.v.Commit?
}

}
