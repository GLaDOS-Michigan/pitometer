// include "../../Impl/Common/SeqIsUniqueDef.i.dfy"

include "Types.dfy"


module ZooKeeper_ZKDatabase {
// import opened Common__SeqIsUniqueDef_i

import opened ZooKeeper_Types

datatype QuorumPeer = QuorumPeer(
    id:      // my id 
    config:  // The cluster configuration. config[id] is my own endpoint 
    peer: // leader or follower
}