include "../../Common/Native/NativeTypes.s.dfy"
include "../../Common/Framework/EnvironmentTCP.s.dfy"

module ZooKeeper_Types {
import opened Native__NativeTypes_s
import opened EnvironmentTCP_s


/*****************************************************************************************
*                                       Hoststep                                         *
*****************************************************************************************/

datatype FollowerStep = 
    SendFollowerInfo | ProcessLeaderInfo | ProcessSync | ProcessSnap | FStutter

datatype LeaderStep = 
    ProcessFollowerInfo | ProcessEpochAck | PrepSync | DoSync | DoSyncSNAP | ProcessAck | LStutter


datatype ZKStep = F(fs: FollowerStep) | L(ls:LeaderStep)


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