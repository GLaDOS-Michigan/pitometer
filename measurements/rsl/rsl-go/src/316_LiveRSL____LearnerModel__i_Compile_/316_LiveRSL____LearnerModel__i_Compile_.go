// Package _316_LiveRSL____LearnerModel__i_Compile
// Dafny module _316_LiveRSL____LearnerModel__i_Compile compiled into Go

package _316_LiveRSL____LearnerModel__i_Compile

import (
  _dafny "dafny"
_System "System_"
_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
_7_Environment__s_Compile "7_Environment__s_Compile_"
_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
_30_Collections____Seqs__i_Compile "30_Collections____Seqs__i_Compile_"
_32_Math____mul__nonlinear__i_Compile "32_Math____mul__nonlinear__i_Compile_"
_34_Math____mul__auto__proofs__i_Compile "34_Math____mul__auto__proofs__i_Compile_"
_36_Math____mul__auto__i_Compile "36_Math____mul__auto__i_Compile_"
_40_Math____mul__i_Compile "40_Math____mul__i_Compile_"
_42_Math____div__nonlinear__i_Compile "42_Math____div__nonlinear__i_Compile_"
_44_Math____mod__auto__proofs__i_Compile "44_Math____mod__auto__proofs__i_Compile_"
_46_Math____mod__auto__i_Compile "46_Math____mod__auto__i_Compile_"
_50_AppStateMachine__i_Compile "50_AppStateMachine__i_Compile_"
_54_Concrete__NodeIdentity__i_Compile "54_Concrete__NodeIdentity__i_Compile_"
_56_LiveRSL____Types__i_Compile "56_LiveRSL____Types__i_Compile_"
_58_Collections____Sets__i_Compile "58_Collections____Sets__i_Compile_"
_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
_101_LiveRSL____Proposer__i_Compile "101_LiveRSL____Proposer__i_Compile_"
_115_LiveRSL____StateMachine__i_Compile "115_LiveRSL____StateMachine__i_Compile_"
_118_Collections____Maps__i_Compile "118_Collections____Maps__i_Compile_"
_120_LiveRSL____Executor__i_Compile "120_LiveRSL____Executor__i_Compile_"
_123_LiveRSL____Learner__i_Compile "123_LiveRSL____Learner__i_Compile_"
_126_LiveRSL____Replica__i_Compile "126_LiveRSL____Replica__i_Compile_"
_135_Logic____Option__i_Compile "135_Logic____Option__i_Compile_"
_138_Native____NativeTypes__i_Compile "138_Native____NativeTypes__i_Compile_"
_141_Libraries____base__s_Compile "141_Libraries____base__s_Compile_"
_143_Math____power2__s_Compile "143_Math____power2__s_Compile_"
_145_Math____power__s_Compile "145_Math____power__s_Compile_"
_149_Math____power__i_Compile "149_Math____power__i_Compile_"
_153_Math____div__def__i_Compile "153_Math____div__def__i_Compile_"
_157_Math____div__boogie__i_Compile "157_Math____div__boogie__i_Compile_"
_162_Math____div__auto__proofs__i_Compile "162_Math____div__auto__proofs__i_Compile_"
_164_Math____div__auto__i_Compile "164_Math____div__auto__i_Compile_"
_166_Math____div__i_Compile "166_Math____div__i_Compile_"
_168_Math____power2__i_Compile "168_Math____power2__i_Compile_"
_170_Common____Util__i_Compile "170_Common____Util__i_Compile_"
_174_Common____MarshallInt__i_Compile "174_Common____MarshallInt__i_Compile_"
_176_Common____GenericMarshalling__i_Compile "176_Common____GenericMarshalling__i_Compile_"
_180_Common____UdpClient__i_Compile "180_Common____UdpClient__i_Compile_"
_182_Common____SeqIsUniqueDef__i_Compile "182_Common____SeqIsUniqueDef__i_Compile_"
_185_Common____SeqIsUnique__i_Compile "185_Common____SeqIsUnique__i_Compile_"
_191_GenericRefinement__i_Compile "191_GenericRefinement__i_Compile_"
_194_Common____NodeIdentity__i_Compile "194_Common____NodeIdentity__i_Compile_"
_197_LiveRSL____AppInterface__i_Compile "197_LiveRSL____AppInterface__i_Compile_"
_214_LiveRSL____CTypes__i_Compile "214_LiveRSL____CTypes__i_Compile_"
_217_LiveRSL____CMessage__i_Compile "217_LiveRSL____CMessage__i_Compile_"
_225_LiveRSL____CMessageRefinements__i_Compile "225_LiveRSL____CMessageRefinements__i_Compile_"
_228_LiveRSL____PacketParsing__i_Compile "228_LiveRSL____PacketParsing__i_Compile_"
_234_Common____UpperBound__i_Compile "234_Common____UpperBound__i_Compile_"
_236_LiveRSL____ParametersState__i_Compile "236_LiveRSL____ParametersState__i_Compile_"
_238_LiveRSL____CPaxosConfiguration__i_Compile "238_LiveRSL____CPaxosConfiguration__i_Compile_"
_240_LiveRSL____ConstantsState__i_Compile "240_LiveRSL____ConstantsState__i_Compile_"
_243_LiveRSL____PaxosWorldState__i_Compile "243_LiveRSL____PaxosWorldState__i_Compile_"
_245_LiveRSL____ReplicaConstantsState__i_Compile "245_LiveRSL____ReplicaConstantsState__i_Compile_"
_251_LiveRSL____ElectionState__i_Compile "251_LiveRSL____ElectionState__i_Compile_"
_254_LiveRSL____ProposerState__i_Compile "254_LiveRSL____ProposerState__i_Compile_"
_265_LiveRSL____COperationNumberSort__i_Compile "265_LiveRSL____COperationNumberSort__i_Compile_"
_267_LiveRSL____CLastCheckpointedMap__i_Compile "267_LiveRSL____CLastCheckpointedMap__i_Compile_"
_269_LiveRSL____AcceptorState__i_Compile "269_LiveRSL____AcceptorState__i_Compile_"
_276_LiveRSL____ExecutorState__i_Compile "276_LiveRSL____ExecutorState__i_Compile_"
_278_LiveRSL____LearnerState__i_Compile "278_LiveRSL____LearnerState__i_Compile_"
_283_LiveRSL____CClockReading__i_Compile "283_LiveRSL____CClockReading__i_Compile_"
_285_LiveRSL____ReplicaState__i_Compile "285_LiveRSL____ReplicaState__i_Compile_"
_294_LiveRSL____MinCQuorumSize__i_Compile "294_LiveRSL____MinCQuorumSize__i_Compile_"
_297_LiveRSL____ElectionModel__i_Compile "297_LiveRSL____ElectionModel__i_Compile_"
_301_Impl____LiveRSL____Broadcast__i_Compile "301_Impl____LiveRSL____Broadcast__i_Compile_"
_305_LiveRSL____ProposerLemmas__i_Compile "305_LiveRSL____ProposerLemmas__i_Compile_"
_308_LiveRSL____ProposerModel__i_Compile "308_LiveRSL____ProposerModel__i_Compile_"
_312_LiveRSL____AcceptorModel__i_Compile "312_LiveRSL____AcceptorModel__i_Compile_"
)
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ _0_Native____NativeTypes__s_Compile.Dummy__
var _ _2_Collections____Maps2__s_Compile.Dummy__
var _ _5_Temporal____Temporal__s_Compile.Dummy__
var _ _7_Environment__s_Compile.Dummy__
var _ _9_Native____Io__s_Compile.Dummy__
var _ _26_Collections____Seqs__s_Compile.Dummy__
var _ _30_Collections____Seqs__i_Compile.Dummy__
var _ _32_Math____mul__nonlinear__i_Compile.Dummy__
var _ _34_Math____mul__auto__proofs__i_Compile.Dummy__
var _ _36_Math____mul__auto__i_Compile.Dummy__
var _ _40_Math____mul__i_Compile.Dummy__
var _ _42_Math____div__nonlinear__i_Compile.Dummy__
var _ _44_Math____mod__auto__proofs__i_Compile.Dummy__
var _ _46_Math____mod__auto__i_Compile.Dummy__
var _ _50_AppStateMachine__i_Compile.Dummy__
var _ _54_Concrete__NodeIdentity__i_Compile.Dummy__
var _ _56_LiveRSL____Types__i_Compile.Dummy__
var _ _58_Collections____Sets__i_Compile.Dummy__
var _ _61_LiveRSL____Configuration__i_Compile.Dummy__
var _ _65_LiveRSL____Message__i_Compile.Dummy__
var _ _68_LiveRSL____Environment__i_Compile.Dummy__
var _ _71_LiveRSL____ClockReading__i_Compile.Dummy__
var _ _74_Common____UpperBound__s_Compile.Dummy__
var _ _76_LiveRSL____Parameters__i_Compile.Dummy__
var _ _78_LiveRSL____Constants__i_Compile.Dummy__
var _ _85_LiveRSL____Broadcast__i_Compile.Dummy__
var _ _91_Collections____CountMatches__i_Compile.Dummy__
var _ _93_LiveRSL____Acceptor__i_Compile.Dummy__
var _ _99_LiveRSL____Election__i_Compile.Dummy__
var _ _101_LiveRSL____Proposer__i_Compile.Dummy__
var _ _115_LiveRSL____StateMachine__i_Compile.Dummy__
var _ _118_Collections____Maps__i_Compile.Dummy__
var _ _120_LiveRSL____Executor__i_Compile.Dummy__
var _ _123_LiveRSL____Learner__i_Compile.Dummy__
var _ _126_LiveRSL____Replica__i_Compile.Dummy__
var _ _135_Logic____Option__i_Compile.Dummy__
var _ _138_Native____NativeTypes__i_Compile.Dummy__
var _ _141_Libraries____base__s_Compile.Dummy__
var _ _143_Math____power2__s_Compile.Dummy__
var _ _145_Math____power__s_Compile.Dummy__
var _ _149_Math____power__i_Compile.Dummy__
var _ _153_Math____div__def__i_Compile.Dummy__
var _ _157_Math____div__boogie__i_Compile.Dummy__
var _ _162_Math____div__auto__proofs__i_Compile.Dummy__
var _ _164_Math____div__auto__i_Compile.Dummy__
var _ _166_Math____div__i_Compile.Dummy__
var _ _168_Math____power2__i_Compile.Dummy__
var _ _170_Common____Util__i_Compile.Dummy__
var _ _174_Common____MarshallInt__i_Compile.Dummy__
var _ _176_Common____GenericMarshalling__i_Compile.Dummy__
var _ _180_Common____UdpClient__i_Compile.Dummy__
var _ _182_Common____SeqIsUniqueDef__i_Compile.Dummy__
var _ _185_Common____SeqIsUnique__i_Compile.Dummy__
var _ _191_GenericRefinement__i_Compile.Dummy__
var _ _194_Common____NodeIdentity__i_Compile.Dummy__
var _ _197_LiveRSL____AppInterface__i_Compile.Dummy__
var _ _214_LiveRSL____CTypes__i_Compile.Dummy__
var _ _217_LiveRSL____CMessage__i_Compile.Dummy__
var _ _225_LiveRSL____CMessageRefinements__i_Compile.Dummy__
var _ _228_LiveRSL____PacketParsing__i_Compile.Dummy__
var _ _234_Common____UpperBound__i_Compile.Dummy__
var _ _236_LiveRSL____ParametersState__i_Compile.Dummy__
var _ _238_LiveRSL____CPaxosConfiguration__i_Compile.Dummy__
var _ _240_LiveRSL____ConstantsState__i_Compile.Dummy__
var _ _243_LiveRSL____PaxosWorldState__i_Compile.Dummy__
var _ _245_LiveRSL____ReplicaConstantsState__i_Compile.Dummy__
var _ _251_LiveRSL____ElectionState__i_Compile.Dummy__
var _ _254_LiveRSL____ProposerState__i_Compile.Dummy__
var _ _265_LiveRSL____COperationNumberSort__i_Compile.Dummy__
var _ _267_LiveRSL____CLastCheckpointedMap__i_Compile.Dummy__
var _ _269_LiveRSL____AcceptorState__i_Compile.Dummy__
var _ _276_LiveRSL____ExecutorState__i_Compile.Dummy__
var _ _278_LiveRSL____LearnerState__i_Compile.Dummy__
var _ _283_LiveRSL____CClockReading__i_Compile.Dummy__
var _ _285_LiveRSL____ReplicaState__i_Compile.Dummy__
var _ _294_LiveRSL____MinCQuorumSize__i_Compile.Dummy__
var _ _297_LiveRSL____ElectionModel__i_Compile.Dummy__
var _ _301_Impl____LiveRSL____Broadcast__i_Compile.Dummy__
var _ _305_LiveRSL____ProposerLemmas__i_Compile.Dummy__
var _ _308_LiveRSL____ProposerModel__i_Compile.Dummy__
var _ _312_LiveRSL____AcceptorModel__i_Compile.Dummy__

type Dummy__ struct{}




// Definition of class Default__
type Default__ struct {
  dummy byte
}

func New_Default___() *Default__ {
  _this := Default__{}

  return &_this
}

type CompanionStruct_Default___ struct {
}
var Companion_Default___ = CompanionStruct_Default___ {
}

func (_this *Default__) Equals(other *Default__) bool {
  return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
  other, ok := x.(*Default__)
return ok && _this.Equals(other)
}

func (*Default__) String() string {
  return "_316_LiveRSL____LearnerModel__i_Compile.Default__"
}

func Type_Default___() _dafny.Type {
  return type_Default___{}
}

type type_Default___ struct {
}

func (_this type_Default___) Default() interface{} {
  return (*Default__)(nil)
}

func (_this type_Default___) String() string {
  return "_316_LiveRSL____LearnerModel__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) LearnerModel__Process2b(learner _278_LiveRSL____LearnerState__i_Compile.CLearnerState, executor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState, packet _217_LiveRSL____CMessage__i_Compile.CPacket) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
  var learner_k _278_LiveRSL____LearnerState__i_Compile.CLearnerState = _278_LiveRSL____LearnerState__i_Compile.Type_CLearnerState_().Default().(_278_LiveRSL____LearnerState__i_Compile.CLearnerState)
  var _ = learner_k
  { }
  { }
  var _4984_msg _217_LiveRSL____CMessage__i_Compile.CMessage
  var _ = _4984_msg
  _4984_msg = (packet).Dtor_msg()
  var _4985_src _9_Native____Io__s_Compile.EndPoint
  var _ = _4985_src
  _4985_src = (packet).Dtor_src()
  var _4986_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
  var _ = _4986_opn
  _4986_opn = (_4984_msg).Dtor_opn__2b()
  var _4987_isBalLt1 bool
  var _ = _4987_isBalLt1
var _out225 bool
  var _ = _out225
_out225 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLt((_4984_msg).Dtor_bal__2b(), (learner).Dtor_max__ballot__seen())
_4987_isBalLt1 = _out225
  var _4988_isBalLt2 bool
  var _ = _4988_isBalLt2
var _out226 bool
  var _ = _out226
_out226 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLt((learner).Dtor_max__ballot__seen(), (_4984_msg).Dtor_bal__2b())
_4988_isBalLt2 = _out226
  var _4989_srcIsReplica bool
  var _ = _4989_srcIsReplica
  _4989_srcIsReplica = (((((learner).Dtor_rcs()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Contains(_4985_src)
  { }
  { }
  { }
  { }
  { }
  { }
  if ((!(_4989_srcIsReplica)) || (_4987_isBalLt1)) {
    learner_k = learner
  } else if (_4988_isBalLt2) {
    var _4990_tup_k _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple
    var _ = _4990_tup_k
    _4990_tup_k = _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple{_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple_CLearnerTuple{_dafny.SeqOf((packet).Dtor_src()), (_4984_msg).Dtor_val__2b()}}
    learner_k = func (_pat_let86_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
      return func (_4991_dt__update__tmp_h0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
        return func (_pat_let87_0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
          return func (_4992_dt__update_hunexecuted__ops_h0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
            return func (_pat_let88_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
              return func (_4993_dt__update_hmax__ballot__seen_h0 _214_LiveRSL____CTypes__i_Compile.CBallot) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
                return _278_LiveRSL____LearnerState__i_Compile.CLearnerState{_278_LiveRSL____LearnerState__i_Compile.CLearnerState_CLearnerState{(_4991_dt__update__tmp_h0).Dtor_rcs(), _4993_dt__update_hmax__ballot__seen_h0, _4992_dt__update_hunexecuted__ops_h0, (_4991_dt__update__tmp_h0).Dtor_sendDecision(), (_4991_dt__update__tmp_h0).Dtor_opn(), (_4991_dt__update__tmp_h0).Dtor_recv2b(), (_4991_dt__update__tmp_h0).Dtor_recvCp()}}
              }(_pat_let88_0)
            }((_4984_msg).Dtor_bal__2b())
          }(_pat_let87_0)
        }(_dafny.NewMapBuilder().Add(_4986_opn, _4990_tup_k).ToMap())
      }(_pat_let86_0)
    }(learner)
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
  } else if (!((learner).Dtor_unexecuted__ops()).Contains(_4986_opn)) {
    { }
    var _4994_tup_k _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple
    var _ = _4994_tup_k
    _4994_tup_k = _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple{_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple_CLearnerTuple{_dafny.SeqOf((packet).Dtor_src()), (_4984_msg).Dtor_val__2b()}}
    learner_k = func (_pat_let89_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
      return func (_4995_dt__update__tmp_h2 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
        return func (_pat_let90_0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
          return func (_4996_dt__update_hunexecuted__ops_h1 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
            return _278_LiveRSL____LearnerState__i_Compile.CLearnerState{_278_LiveRSL____LearnerState__i_Compile.CLearnerState_CLearnerState{(_4995_dt__update__tmp_h2).Dtor_rcs(), (_4995_dt__update__tmp_h2).Dtor_max__ballot__seen(), _4996_dt__update_hunexecuted__ops_h1, (_4995_dt__update__tmp_h2).Dtor_sendDecision(), (_4995_dt__update__tmp_h2).Dtor_opn(), (_4995_dt__update__tmp_h2).Dtor_recv2b(), (_4995_dt__update__tmp_h2).Dtor_recvCp()}}
          }(_pat_let90_0)
        }(((learner).Dtor_unexecuted__ops()).Update(_4986_opn, _4994_tup_k))
      }(_pat_let89_0)
    }(learner)
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
  } else if (((((learner).Dtor_unexecuted__ops()).Get(_4986_opn).(_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple)).Dtor_received__2b__message__senders()).Contains((packet).Dtor_src())) {
    learner_k = learner
  } else {
    var _4997_tup _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple
    var _ = _4997_tup
    _4997_tup = ((learner).Dtor_unexecuted__ops()).Get(_4986_opn).(_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple)
    var _4998_tup_k _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple
    var _ = _4998_tup_k
    _4998_tup_k = func (_pat_let91_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple) _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple {
      return func (_4999_dt__update__tmp_h4 _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple) _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple {
        return func (_pat_let92_0 _dafny.Seq) _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple {
          return func (_5000_dt__update_hreceived__2b__message__senders_h0 _dafny.Seq) _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple {
            return _278_LiveRSL____LearnerState__i_Compile.CLearnerTuple{_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple_CLearnerTuple{_5000_dt__update_hreceived__2b__message__senders_h0, (_4999_dt__update__tmp_h4).Dtor_candidate__learned__value()}}
          }(_pat_let92_0)
        }(((_4997_tup).Dtor_received__2b__message__senders()).Concat(_dafny.SeqOf((packet).Dtor_src())))
      }(_pat_let91_0)
    }(_4997_tup)
    learner_k = func (_pat_let93_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
      return func (_5001_dt__update__tmp_h5 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
        return func (_pat_let94_0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
          return func (_5002_dt__update_hunexecuted__ops_h2 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
            return _278_LiveRSL____LearnerState__i_Compile.CLearnerState{_278_LiveRSL____LearnerState__i_Compile.CLearnerState_CLearnerState{(_5001_dt__update__tmp_h5).Dtor_rcs(), (_5001_dt__update__tmp_h5).Dtor_max__ballot__seen(), _5002_dt__update_hunexecuted__ops_h2, (_5001_dt__update__tmp_h5).Dtor_sendDecision(), (_5001_dt__update__tmp_h5).Dtor_opn(), (_5001_dt__update__tmp_h5).Dtor_recv2b(), (_5001_dt__update__tmp_h5).Dtor_recvCp()}}
          }(_pat_let94_0)
        }(((learner).Dtor_unexecuted__ops()).Update(_4986_opn, _4998_tup_k))
      }(_pat_let93_0)
    }(learner)
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
  }
  return learner_k
}
func (_this *CompanionStruct_Default___) LearnerModel__ForgetDecision(learner _278_LiveRSL____LearnerState__i_Compile.CLearnerState, opn _214_LiveRSL____CTypes__i_Compile.COperationNumber) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
  goto TAIL_CALL_START
TAIL_CALL_START:
var learner_k _278_LiveRSL____LearnerState__i_Compile.CLearnerState = _278_LiveRSL____LearnerState__i_Compile.Type_CLearnerState_().Default().(_278_LiveRSL____LearnerState__i_Compile.CLearnerState)
  var _ = learner_k
  { }
  if (((learner).Dtor_unexecuted__ops()).Contains(opn)) {
    learner_k = func (_pat_let95_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
      return func (_5003_dt__update__tmp_h0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
        return func (_pat_let96_0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
          return func (_5004_dt__update_hunexecuted__ops_h0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
            return _278_LiveRSL____LearnerState__i_Compile.CLearnerState{_278_LiveRSL____LearnerState__i_Compile.CLearnerState_CLearnerState{(_5003_dt__update__tmp_h0).Dtor_rcs(), (_5003_dt__update__tmp_h0).Dtor_max__ballot__seen(), _5004_dt__update_hunexecuted__ops_h0, (_5003_dt__update__tmp_h0).Dtor_sendDecision(), (_5003_dt__update__tmp_h0).Dtor_opn(), (_5003_dt__update__tmp_h0).Dtor_recv2b(), (_5003_dt__update__tmp_h0).Dtor_recvCp()}}
          }(_pat_let96_0)
        }(_118_Collections____Maps__i_Compile.Companion_Default___.RemoveElt((learner).Dtor_unexecuted__ops(), opn))
      }(_pat_let95_0)
    }(learner)
    { }
    { }
  } else {
    learner_k = learner
  }
  return learner_k
}
func (_this *CompanionStruct_Default___) LearnerModel__ForgetOperationsBefore(learner _278_LiveRSL____LearnerState__i_Compile.CLearnerState, ops__complete _214_LiveRSL____CTypes__i_Compile.COperationNumber) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
  goto TAIL_CALL_START
TAIL_CALL_START:
var learner_k _278_LiveRSL____LearnerState__i_Compile.CLearnerState = _278_LiveRSL____LearnerState__i_Compile.Type_CLearnerState_().Default().(_278_LiveRSL____LearnerState__i_Compile.CLearnerState)
  var _ = learner_k
  var _5005_unexecuted__ops_k _dafny.Map
  var _ = _5005_unexecuted__ops_k
  _5005_unexecuted__ops_k = func () _dafny.Map {
    var _coll5 = _dafny.NewMapBuilder()
    var _ = _coll5
for _iter17 := _dafny.Iterate(((learner).Dtor_unexecuted__ops()).Keys().Elements());; {
      _val17, _ok17 := _iter17()
if !_ok17 { break }
_5006_op := _val17.(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
if ((((learner).Dtor_unexecuted__ops()).Contains(_5006_op)) && (((_5006_op).Dtor_n()) >= ((ops__complete).Dtor_n()))) {
        _coll5.Add(_5006_op,((learner).Dtor_unexecuted__ops()).Get(_5006_op).(_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple))
      }
    }
    return _coll5.ToMap()
  }()
  learner_k = func (_pat_let97_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
    return func (_5007_dt__update__tmp_h0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
      return func (_pat_let98_0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
        return func (_5008_dt__update_hunexecuted__ops_h0 _dafny.Map) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
          return _278_LiveRSL____LearnerState__i_Compile.CLearnerState{_278_LiveRSL____LearnerState__i_Compile.CLearnerState_CLearnerState{(_5007_dt__update__tmp_h0).Dtor_rcs(), (_5007_dt__update__tmp_h0).Dtor_max__ballot__seen(), _5008_dt__update_hunexecuted__ops_h0, (_5007_dt__update__tmp_h0).Dtor_sendDecision(), (_5007_dt__update__tmp_h0).Dtor_opn(), (_5007_dt__update__tmp_h0).Dtor_recv2b(), (_5007_dt__update__tmp_h0).Dtor_recvCp()}}
        }(_pat_let98_0)
      }(_5005_unexecuted__ops_k)
    }(_pat_let97_0)
  }(learner)
  { }
  return learner_k
}
// End of class Default__
