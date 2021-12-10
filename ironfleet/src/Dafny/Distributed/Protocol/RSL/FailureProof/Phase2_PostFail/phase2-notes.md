## Old Invariant

* `P2Assumption`
    * `RslConsistency`
        * `ConstantsAllConsistentInv`
        * `WellFormedLConfiguration`
    * `LeaderAlwaysZero`
    * `BoundedQueueingAssumption`
        * for each replica performing a receive action, we have `s.t_replicas[idx].ts <= io.r.msg.ts + MaxQueueTime`
* `Phase2Invariant`
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


## Flow of Client Request in Phase 2

### Leader Receiving a Request from Client

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
    * I saw that acceptors have previously accepted some value for this slot, by looking into my set of 1b messages I received for this view. Then `LProposerNominateOldValueAndSend2a`.
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
3. Add `(bal, val)` to my `votes` map for this slot.
4. Acceptor broadcasts 2b message to everyone, updates promised ballot `max_bal`

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

### Leader Starting His Reign

1. Server runs `LSchedulerNext`, of the `LReplicaNoReceiveNext` variety.
2. The action of interest is number 1: `LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a`
3. If the new election indicates that I should become the leader, `s.election_state.current_view.proposer_id == s.constants.my_index`, then 
    * Enter Phase 1
    * Set `max_ballot_i_sent_1a` to this new view
    * Set `received_1b_packets` to the empty set
    * Set `highest_seqno_requested_by_client_this_view` to the empty map
    * Update `request_queue`
    * Broadcast new 1a message to everyone

### Server Receiving a 1a Message

1. Server runs `LSchedulerNext`, and enters the sub-clause `LReplicaNextProcess1a`
2. If this 1a message has a ballot strictly larger than what I've seen, then send a 1b response containing all of my votes.

### Leader Discovering Accepted Values in 1b Messages

1. Server runs `LSchedulerNext`, and enters the sub-clause `LReplicaNextProcess1b`
2. If I am in Phase 1, and the 1b packet matches the 1a ballot I sent `max_ballot_i_sent_1a`, and this 1b is not a duplicate, then run `LProposerProcess1b`.
3. This adds the source of the 1b packet to my `received_1b_packets` set.
4. Next, run `LAcceptorTruncateLog`

### Leader Entering Phase 2

1. Server runs `LSchedulerNext`, of the `LReplicaNoReceiveNext` variety.
2. The action of interest is number 2: `LReplicaNextSpontaneousMaybeEnterPhase2`
3. If I amass a quorum of 1b messages, then
    * Enter Phase 2
    * Broadcast `RslMessage_StartingPhase2` message to everyone

1. Servers receiving this message run `LReplicaNextProcessStartingPhase2`
2. They broadcast `RslMessage_AppStateRequest` to everyone
3. Servers receiving this state request runs `LReplicaNextProcessAppStateRequest`
4. They respond with `RslMessage_AppStateSupply`
5. Servers receiving this state suppy runs `LExecutorProcessAppStateSupply`


## Failure Detection

### Server Sending and Receiving Heartbeats

1. Server runs `LSchedulerNext`, of the `LReplicaNoReceiveNext` variety.
2. The action of interest is number 9: `LReplicaNextReadClockMaybeSendHeartbeat`
3. Heartbeat contains
    * `election_state.current_view`
    * Boolean indicating if I suspect the current view
    * `s.executor.ops_complete`

1. Server runs `LSchedulerNext`, and enters the sub-clause `LReplicaNextProcessHeartbeat`
2. First, server runs `LProposerProcessHeartbeat`. This runs `ElectionStateProcessHeartbeat` to update election state
    * If heartbeat indicates that sender suspects the current view, add the source to `es.current_view_suspectors`
    * If heartbeat indicates that sender is in a larger view, then update my current view to match
3. If current view is increased, then set `current_state == 0` to downgrade to follower status, and reset `request_queue`
4. Next, server runs `LAcceptorProcessHeartbeat`. This does some checkpointing stuff

### Server Triggering View Change
     
1. Server runs `LSchedulerNext`, of the `LReplicaNoReceiveNext` variety.
2. The action of interest is number 8: `LReplicaNextReadClockCheckForQuorumOfViewSuspicions`.
3. This runs `ElectionStateCheckForQuorumOfViewSuspicions`. If I amassed a quorum of suspectors, then 
    * Increment view to `ComputeSuccessorView(es.current_view, es.constants.all)`
    * Set `current_view_suspectors` to empty
    * Update `requests_received_prev_epochs` and `requests_received_this_epoch`
4. If current view is increased, then set `current_state == 0` to downgrade to follower status, and reset `request_queue`


At what point will I suspect the current leader?

## New Initial State

* Leader's set of 1b packets have the original client request to be proposed
* Every packet at genesis has some initial time bound. By simple rules, no need such bound for nodes
* Every 2a packet has ballot (1, 1) or (1, 0).
* Leader has election state and view (1, 1).
* No response to client in the system.


## Goal: First Figure Out TimeBound2aDelivery


## Notes

### Resolved

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
    * That, *and* it can maybe see it accepted by some acceptor through phase 1.

### Pending

* What happens when I get a timeout `LIoOpTimeoutReceive`?
* acceptor and nodes highest ballot is (1, 1) should be some kind of invariant?
* May want to say that `Old_1bTS` is `Old_1aTS + proc_1a + D` 
* May want to say that `Old_2bTS` is `Old_2aTS + proc_2a + D` 



## Misc 

/home/nudzhang/Documents/pitometer/ironfleet/src/Dafny/Distributed/Protocol/RSL/FailureProof


```
var idx, ios :| TimestampedRslNextOneReplica(s, s', idx, ios);
var us, us', uios := UntimestampRslState(s), UntimestampRslState(s'), UntagLIoOpSeq(ios);
var nextActionIndex := us.replicas[idx].nextActionIndex;

if nextActionIndex == 0 {
    assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
} else {
    var r, r' := us.replicas[idx], us'.replicas[idx];
    if nextActionIndex == 1 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    } else if nextActionIndex == 2 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    } else if nextActionIndex == 3 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    }else if nextActionIndex == 4 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    }else if nextActionIndex == 5 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    }else if nextActionIndex == 6 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    }else if nextActionIndex == 7 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    }else if nextActionIndex == 8 {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    }else {
        assert BoundaryConditionInvariant_ExistingPacketsBallot(s');
    }
}




var lr, lr' := ts.t_replicas[1].v.replica, ts'.t_replicas[1].v.replica;
assert All2aPackets_BalLeq_Opn(ts', Ballot(1, 1), opn);
assert All2bPackets_BalLeq_Opn(ts', Ballot(1, 1), opn);
assert (exists pkt :: pkt in ts'.t_environment.sentPackets && IsNew2bPacket(pkt, opn));
assert PerformanceGuarantee_2a(ts', opn);
assert PerformanceGuarantee_2b(ts', opn);
assert PerformanceGuarantee_Response(ts');
assert 0 <= ts'.t_replicas[1].v.nextActionIndex <= 9;
assert lr'.proposer.current_state == 2;
assert lr'.proposer.next_operation_number_to_propose > opn;

// Learner and Executor states
assert BalLeq(lr'.learner.max_ballot_seen, Ballot(1, 1));

assert (Get2bCount(lr', opn, Ballot(1, 1)) < LMinQuorumSize(ts.constants.config)
==> && lr'.executor.next_op_to_execute.OutstandingOpUnknown?
    && lr'.executor.ops_complete == opn
    && (!exists pkt :: pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt))
);

assert (Get2bCount(lr', opn, Ballot(1, 1)) == LMinQuorumSize(ts'.constants.config)
    ==> TimeLe(ts'.t_replicas[1].ts, TimeBoundPhase2LeaderPost(nextActionIndex')));

assert (Get2bCount(lr', opn, Ballot(1, 1)) == LMinQuorumSize(ts'.constants.config)
    ==> && (nextActionIndex' < 6 ==> 
            && lr'.executor.ops_complete == opn 
            && lr'.executor.next_op_to_execute.OutstandingOpUnknown?
            && (!exists pkt :: pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt))
        )
        && (nextActionIndex' == 6 ==> 
            && lr'.executor.ops_complete == opn 
            && lr'.executor.next_op_to_execute.OutstandingOpKnown?
            && (!exists pkt :: pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt))
        )
        && (nextActionIndex' > 6 ==> 
            && lr'.executor.ops_complete > opn
            && (exists pkt :: pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt))
        )
    );

assert ((exists pkt :: pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt))
==> && lr'.executor.ops_complete > opn);

assert (forall opn' | opn' in lr'.learner.unexecuted_learner_state :: opn' == opn);

assert (opn in lr'.learner.unexecuted_learner_state 
    ==>
    && lr'.learner.max_ballot_seen == Ballot(1, 1)
    && (forall id :: id in lr'.learner.unexecuted_learner_state[opn].received_2b_message_senders ==> id in ts'.constants.config.replica_ids)
);

assert After_2b_Sent_Invariant(ts', opn);
```