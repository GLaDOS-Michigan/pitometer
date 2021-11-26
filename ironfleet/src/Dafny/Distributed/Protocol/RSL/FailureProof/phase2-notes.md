## Preconditions

* `RslAssumption`
    * `RslConsistency`
        * `ConstantsAllConsistentInv`
        * `WellFormedLConfiguration`
    * `LeaderAlwaysZero`
    * `BoundedQueueingAssumption`
        * for each replica performing a receive action, we have `s.t_replicas[idx].ts <= io.r.msg.ts + MaxQueueTime`
* `RslPerfInvariant`
    * `GenericPhase2Invariant`
        * `AlwaysInvariant`
            * ServersAreNotClients
            * All undelivered 2a and 2b packets is sent by a client
            * Every unexecuted learned value is sent by a client
            * Every op to execute is sent by a client
        * map `progresses` is valid
        * `s.t_replicas[0]` proposer has an empty `request_queue`
        * `GenericPhase2UndeliveredPacketInvariant`
            * Every undelivered packet is a 2a or 2b packet
            * Every 2a packet
                * has dest in `progresses`
                * has leader as the source
                * has the specified `opn` number
                * has Ballot (1,0)
                * has `TimeLe(pkt.msg.ts, TimeBound2aDelivery(req_time))`
                * if self delivery, then has `(pkt.msg.ts, TimeBound2aSelfDelivery(req_time))`
                * has timestamp `t2a + minD < pkt.msg.ts <= t2a + D`
                * is the only 2a packet for its destination
            * Every 2b packet
                * if has dst as replica 0, then 
                    * `progresses[src]=p2b`
                    * has `TimeLe(pkt.msg.ts, TimeBound2bDelivery(req_time))`
                    * is the only 2b packet to the leader from this source
                * has the specified `opn` number
                * has Ballot (1,0)
                * has `pkt.msg.ts > t2a + 2*minD`
            * `Progress2aProperty`
                * every `P2a` entry in `progresses` implies there is a undelivered 2a packet
        * For every id such that `progresses[id]` is P2a, its learner state is an empty map
    * `Phase2AcceptedLeaderInvariant`
        * `s.t_replicas[0]` proposer has an empty `request_queue`
        * Leader's `nextActionIndex` is in the range `[0, 9]`
        * Some bounds on `opn` number
        * 2b message count less than or equal to one quorum
        * If 2b message count is strictly less than a quorum, then `s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpUnknown?`
        * If next action index of leader is 0 or >= 7, then 2b is not a quorum
        * If 2b message is of quorum size, then leader time bound is    `TimeBoundPhase2Leader(TimeBound2bDelivery(req_time), s.t_replicas[0].v.nextActionIndex))`
        * If leader next action index is 6 and 2b is a quorum, then `s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpKnown?`
        * If `s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpKnown?`, then 2b is of quorum size
        * `ProposedLeaderNextOperationNumberInvariant`
        * Leader is not in P2a in progresses
        * Every unexecuted operation in leader replica is `opn`
        * Leader's max ballot seem is less than `(1, 0)`
        * If `opn` is in `s.t_replicas[0].v.replica.learner.unexecuted_learner_state`, then 
            * Leader's max ballot seen is `(1, 0)`
            * All senders of 2b mesages, `id`, has `progresses[id] == P2done`
    * `Phase2UnacceptedLeaderInvariant`
        * `s.t_replicas[0]` proposer has an empty `request_queue`
        * Leader's `nextActionIndex` is in the range `[0, 9]`
        * Some bounds on `opn` number
        * `s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpUnknown?`
        * Leader's max ballot seem is less than `(1, 0)`
        * Leader is of state P2a in `progresses`
        * Leader's `unexecuted_learner_state` is empty map
        * Leader's next op to execute is `OutstandingOpUnknown()`
        * `ProposedLeaderNextOperationNumberInvariant`


## Notes

* What ballot comes after the initial Ballot `(0, 1)`?
    * Ballot is of type `Ballot(seqno:int, proposer_id:int)`
    * The function `ComputeSuccessorView` increases the `proposer_id` field if the resulting number is a valid replica. Else, it increases `seqno` while resetting `proposer_id` to zero.
    * Hence, the Ballot after `(0, 1)` is `(1, 1)`.
* What is `proposer.request_queue`?
    * As defined in proposer, it is "Values that clients have requested that I need to eventually
    propose, in the order I should propose them"
* Why is it part of the old invariant that the leader always has an empty `request_queue`? Shouldn't it get filled up when it gets a client request?
    * `request_queue` gets appended to when a new client request is received.
    * When are items removed from the queue? When the leader sends 2a messages
    * Not sure why we can say that it is always empty though. 


## Flow of Client Request

### Leader Receiving a Request

1. Client sends request to server 0
2. Replica runs `LSchedulerNext`, and the sub-clause `LReplicaNextProcessPacket`
3. If packet is not a heartbeat, it runs `LReplicaNextProcessPacketWithoutReadingClock`
4. Replica runs `LReplicaNextProcessRequest` for a client request.
5. If request is already in reply cache, then executor processes the request `LExecutorProcessRequest`. Otherwise, run `LProposerProcessRequest`
6. If I am the leader, update `election_state`, `request_queue`, and `highest_seqno_requested_by_client_this_view`
7. Now the request is part of the request queue, ready to be proposed.

### Leader Proposing a Request, and a Server Receiving it

1. 


* Who does client send requests to?
    * Client sends request to server 0. When server 0 fails, it tries server 1.
* The question becomes, after server 0 fails, how does server 1 learn of a pending request? Is it solely through client re-transmission?


|s.t_replicas[0].v.replica.proposer.request_queue| == 0