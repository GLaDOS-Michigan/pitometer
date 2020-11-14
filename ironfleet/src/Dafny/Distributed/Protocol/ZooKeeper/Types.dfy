include "../../Common/Native/NativeTypes.s.dfy"
include "../../Common/Framework/EnvironmentTCP.s.dfy"

module ZooKeeper_Types {
import opened Native__NativeTypes_s
import opened EnvironmentTCP_s

/*****************************************************************************************
*                                      Networking                                        *
*****************************************************************************************/
datatype EndPoint = EndPoint(addr:seq<byte>, port:uint16)
    // UdpPacket_ctor has silly name to ferret out backwards calls

type Config = seq<EndPoint>

datatype ZKMessage = 
    | FollowerInfo(sid:nat, latestZxid:Zxid)
    | LeaderInfo(sid:nat, newZxid:Zxid)
    | AckEpoch(sid:nat, lastLoggedZxid:Zxid, lastAcceptedEpoch:int)

type ZKEnvironment = LEnvironment<EndPoint, ZKMessage>
type ZKPacket = LPacket<EndPoint, ZKMessage>
type ZKIo = LIoOp<EndPoint, ZKMessage>


/*****************************************************************************************
*                                         ZXID                                           *
*****************************************************************************************/

datatype Zxid = Zxid(epoch:int, counter:int)

const NullZxid := Zxid(-1, -1);

predicate ZxidLt(z1:Zxid, z2:Zxid) {
    if z1.epoch < z2.epoch then true
    else z1.epoch == z2.epoch && z1.counter < z2.counter 
}
}