include "../../Common/Framework/Environment.s.dfy"

include "Types.dfy"
include "FDEnvironment.dfy"
include "Node.dfy"
include "Detector.dfy"

module FailureDetector_DistributedSystem {
import opened FailureDetector_Types
import opened FailureDetector_Environment
import opened Environment_s
import opened FailureDetector_Node
import opened FailureDetector_Detector


datatype Agent = N(n:Node) | D(d:Detector)

datatype FD_State = FD_State(
    environment: FDEnvironment,
    servers: map<EndPoint, Agent>
)


predicate FD_Init(config:Config, s:FD_State, heartbeatInterval:int, timeoutInterval:int) {
    // assume time starts at exactly 0 for simplicity
    && LEnvironment_Init(s.environment)
    && s.environment.time == 0  

    // assume configuration
    && ConfigIsUnique(config)
    && config.nodeEp in s.servers && config.detectorEp in s.servers
    && |s.servers| == 2
    && s.servers[config.nodeEp].N?
    && s.servers[config.detectorEp].D?

    // assume initial states
    && NodeInit(s.servers[config.nodeEp].n, heartbeatInterval, config)
    && DetectorInit(s.servers[config.detectorEp].d, timeoutInterval, config)
}


predicate FD_NextOneServer(s:FD_State, s':FD_State, actor:EndPoint, ios:seq<FDIo>)
    requires actor in s.servers
{
    && actor in s'.servers
    && s'.servers == s.servers[actor := s'.servers[actor]]
    &&  if s.servers[actor].N? then
            && s.environment.nextStep == LEnvStepHostIos(actor, ios, NodeStep)
            && s'.servers[actor].N?
            && NodeNext(s.servers[actor].n, s'.servers[actor].n, ios)
        else
            && s.environment.nextStep == LEnvStepHostIos(actor, ios, DetectorStep(s.servers[actor].d.nextActionIndex))
            && s'.servers[actor].D?
            && DetectorNext(s.servers[actor].d, s'.servers[actor].d, ios)
}

predicate FD_Next(s:FD_State, s':FD_State){
    && LEnvironment_Next(s.environment, s'.environment)
    && (exists ep, ios :: ep in s.servers && FD_NextOneServer(s, s', ep, ios))
}
}
