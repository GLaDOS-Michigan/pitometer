include "../../Common/Framework/Environment.s.dfy"
include "../../Common/Native/Io.s.dfy"

module Types_i {
import opened Environment_s
import opened Native__Io_s

datatype LockMessage = Transfer(transfer_epoch:int) | Locked(locked_epoch:int) | Invalid
datatype LockStep = GrantStep() | AcceptStep()

type LockEnvironment = LEnvironment<EndPoint, LockMessage, LockStep>
type LockPacket = LPacket<EndPoint, LockMessage>
type LockIo = LIoOp<EndPoint, LockMessage>

}
