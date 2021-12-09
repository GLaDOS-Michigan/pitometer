# Failure Detection

## Assumptions

* `RslConsistency(s)`
* `BoundedQueueingAssumption(s)`
* `NoStateTransfer(s)`: State-Request, State-Supply and Starting-Phase2 messages are never received.
* `NoExternalSteps(s)`
* `OneAndOnlyOneRequest(s)`: there is a ghost constant `req`, and the system contains only this request.

* `ClockAssumption(s, s')`
    * When a replica `r` reads the clock in `s` to obtain value `t`, it must be that `r.ts <= t <= r'.ts`.



