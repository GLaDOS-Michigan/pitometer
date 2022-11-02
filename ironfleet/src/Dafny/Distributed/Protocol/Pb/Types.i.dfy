include "../../Common/Framework/Environment.s.dfy"
include "../../Common/Native/Io.s.dfy"

module Types_i {
import opened Environment_s
import opened Native__Io_s

datatype PbMessage = Request() | Ack() | ClientReply() | Invalid
datatype PbStep = PrimaryReqStep() | PrimaryRecvStep() | BackupRecvStep()

type PbEnvironment = LEnvironment<EndPoint, PbMessage, PbStep>
type PbPacket = LPacket<EndPoint, PbMessage>
type PbIo = LIoOp<EndPoint, PbMessage>

}
