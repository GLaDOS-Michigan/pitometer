include "../../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "../Types.dfy"
include "../DistributedSystem.dfy"
include "../ZKEnvironment.dfy"
include "../Leader.dfy"


module ZooKeeper_CorrectnessProof {
    import opened ZooKeeper_Types
    import opened ZooKeeper_DistributedSystem
    import opened ZooKeeper_Environment
    import opened ZooKeeper_Leader



predicate IsValidConfig(config:Config, ls:LS_State, f:int) {
    && f >= 1
    && |config| == 2*f + 1 > 0
    && forall ep | ep in config :: ep in ls.servers
}

predicate ExactlyOneLeader(config:Config, ls:LS_State, f:int) {
    IsValidConfig(config, ls, f) ==> (
        && ls.servers[config[0]].LeaderPeer?
        && forall i | 1 <= i < |config| :: ls.servers[config[i]].FollowerPeer?
    )
}


predicate FollowersShareLeaderDB(config:Config, ls:LS_State, f:int) 
    requires IsValidConfig(config, ls, f)
    requires ExactlyOneLeader(config, ls, f)
{
    && var leader := ls.servers[config[0]].leader;
    && if leader.state == L_RUNNING 
        then forall p | p in leader.globals.ackSet :: (
            && 0 <= p < |config|
            && ls.servers[config[p]].FollowerPeer? ==> ls.servers[config[p]].follower.zkdb.commitLog == leader.globals.zkdb.commitLog
        ) else true
}

predicate Safety(config:Config, ls:LS_State, f:int) {
    && IsValidConfig(config, ls, f)
    && ExactlyOneLeader(config, ls, f)
    && FollowersShareLeaderDB(config, ls, f)
}

predicate LeaderDBConstant(config:Config, ls:LS_State, ls':LS_State, f:int) 
    requires LS_Next(ls, ls')
    requires IsValidConfig(config, ls, f) && IsValidConfig(config, ls', f)
    requires ExactlyOneLeader(config, ls, f) && ExactlyOneLeader(config, ls', f)
{
    var l, l' := ls.servers[config[0]].leader, ls'.servers[config[0]].leader;
    l.globals.zkdb.commitLog == l'.globals.zkdb.commitLog
}


lemma Main(config:Config, lb:seq<LS_State>, f:int) 
    requires f >= 1
    requires |lb| > 0
    requires |config| == 2*f + 1
    requires LS_Init(config, lb[0], f)
    requires forall i {:trigger LS_Next(lb[i], lb[i+1])} :: 0 <= i < |lb| - 1 ==> LS_Next(lb[i], lb[i+1]);
    ensures forall i {:trigger LS_Next(lb[i], lb[i+1])} :: 0 <= i < |lb| ==> Safety(config, lb[i], f);
    ensures forall i :: 0 <= i < |lb| - 1 ==> LeaderDBConstant(config, lb[i], lb[i+1], f)
{
    assume false;
}

}