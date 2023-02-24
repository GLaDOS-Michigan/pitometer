include "../../Common/Native/NativeTypes.s.dfy"
include "../../Common/Framework/EnvironmentTCP.s.dfy"
include "ZKDatabase.dfy"
include "Types.dfy"

module ZooKeeper_Environment {
import opened Native__NativeTypes_s
import opened EnvironmentTCP_s
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Types


datatype EndPoint = EndPoint(addr:seq<byte>, port:uint16)
    // UdpPacket_ctor has silly name to ferret out backwards calls

type Config = seq<EndPoint>

datatype ZKMessage = 
    | FollowerInfo(sid:nat, latestZxid:Zxid)
    | LeaderInfo(sid:nat, serial:nat, newZxid:Zxid)
    | AckEpoch(sid:nat, serial:nat, lastLoggedZxid:Zxid, lastAcceptedEpoch:int)
    | SyncDIFF(sid:nat, serial:nat, lastProcessedZxid:Zxid)
    | SyncSNAP(sid:nat, serial:nat, leaderDb: ZKDatabase, lastProcessedZxid:Zxid)
    | SyncTRUNC(sid:nat, serial:nat, lastProcessedZxid:Zxid)
    | NewLeader(sid:nat, serial:nat, newLeaderZxid:Zxid)
    | Commit(sid:nat, serial:nat, txn:Zxid)
    | Ack(sid:nat, serial:nat, ackZxid:Zxid)
    | UpToDate(sid:nat)

type ZKEnvironment = LEnvironment<EndPoint, ZKMessage>
type ZKPacket = LPacket<EndPoint, ZKMessage>
type ZKIo = LIoOp<EndPoint, ZKMessage>

}