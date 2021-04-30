include "../../Common/Native/NativeTypes.s.dfy"
include "../../Common/Framework/Environment.s.dfy"

module FailureDetector_Types {
import opened Native__NativeTypes_s
import opened Environment_s


/*****************************************************************************************
*                                       Hoststep                                         *
*****************************************************************************************/

datatype RslStep = NodeStep | DetectorStep(actionIndex:int)

datatype ClockReading = ClockReading(t:int)

}