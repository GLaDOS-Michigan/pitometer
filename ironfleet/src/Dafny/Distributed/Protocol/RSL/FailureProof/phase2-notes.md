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


* Who does client send requests to?
    * Client sends request to server 0. When server 0 fails, it tries server 1.
* The question becomes, after server 0 fails, how does server 1 learn of a pending request? Is it solely through client re-transmission?


## Flow of Client Request in Phase 2

### Leader Receiving a Request

1. Client sends request to server 0
2. Leader runs `LSchedulerNext`, and the sub-clause `LReplicaNextProcessPacket`
3. If packet is not a heartbeat, it runs `LReplicaNextProcessPacketWithoutReadingClock`
4. Replica runs `LReplicaNextProcessRequest` for a client request.
5. If request is already in reply cache, then executor processes the request `LExecutorProcessRequest`. Otherwise, run `LProposerProcessRequest`
6. If I am the leader, update `election_state`, `request_queue`, and `highest_seqno_requested_by_client_this_view`
7. Now the request is part of the request queue, ready to be proposed.

### Leader Proposing a Request as 2a Message

1. Leader runs `LSchedulerNext`, of the `LReplicaNoReceiveNext` variety.
2. The action of interest is number 3: `LReplicaNextReadClockMaybeNominateValueAndSend2a`
3. There are two paths to take here:
    * I saw that acceptors have previously accepted some value for this slot. Then `LProposerNominateOldValueAndSend2a`.
    * Else, nothing could have been chosen for this slot. Then `LProposerNominateNewValueAndSend2a`.
4. For the purposes of the failure proof, the longest path would be that the old leader got the value chosen, learned the value, and failed just before responding to the client. So the new leader will see the chosen value. 
5. Let's explore the `LProposerNominateOldValueAndSend2a` path.
    * Leader simply broadcasts the highest numbered proposal as a 2a message to everyone. 
6. What about `LProposerNominateNewValueAndSend2a`? 
    * Each proposed value is a batch of requests that is a prefix of `request_queue`
    * `request_queue` gets truncated as appropriate
    * Leader broadcasts new proposal as a 2a message to everyone. 

### Server Receiving a 2a Message

1. Server runs `LSchedulerNext`, and enters the sub-clause `LReplicaNextProcess2a`
2. If 2a message has ballot at least as large as my last promised ballot, then run `LAcceptorProcess2a`
3. Acceptor broadcasts 2b message to everyone, updates promised ballot `max_bal`

### Server Learning a Request

1. Server runs `LSchedulerNext`, and enters the sub-clause `LReplicaNextProcess2b`
2. If this 2b message is for a slot that I have yet learned, then run `LLearnerProcess2b`
3. Note that learner state consists of a `unexecuted_learner_state` map. This maps slot numbers to the pair
    * set of acceptors that accepted this ballot
    * value accepted for the slot with the current ballot
4. `LLearnerProcess2b` has a few cases depending on the ballot of the received 2b message
    * If packet ballot is stale, then ignore the packet.
    * If packet ballot is larger than the largest I've seen, then update `max_ballot_seen`, and set `unexecuted_learner_state` for this slot to `LearnerTuple({packet.src}, m.val_2b)`
    * If packet ballot is current, and is the first I've seen for this slot, set `unexecuted_learner_state` for this slot to `LearnerTuple({packet.src}, m.val_2b)`.
    * If packet ballot is current, but this is a duplicate from a previously seen acceptor, then ignore.
    * If packet ballot is current, and from a new source for this slot, then add the new source to the set of sources tracked in `unexecuted_learner_state` for this slot. 

### Server Deciding Executing a Request

1. Server runs `LSchedulerNext`, of the `LReplicaNoReceiveNext` variety.
2. The first action of interest is number 5: `LReplicaNextSpontaneousMaybeMakeDecision`
3. Examines the next undecided slot `s.executor.ops_complete` in `s.learner.unexecuted_learner_state`
4. If this slot has amassed a quorum of 2b replies, then run `LExecutorGetDecision`
5. Server sets the next operation to execute to this slot's value `next_op_to_execute := OutstandingOpKnown(v, bal)`

1. The second action of interest is number 6: `LReplicaNextSpontaneousMaybeExecute`
2. If next operation to execute is decided, `s.executor.next_op_to_execute.OutstandingOpKnown?`, then run `LExecutorExecute`.
3. Run `HandleRequestBatch` to generate new state and client responses
4. Set application state to new state, send replies to client, and update reply cache
5. Reset next slot to execute `s'.next_op_to_execute == OutstandingOpUnknown()`


## Flow of Client Request in Phase 1

### Server Becoming a Leader and Learning of Maybe Chosen Values
