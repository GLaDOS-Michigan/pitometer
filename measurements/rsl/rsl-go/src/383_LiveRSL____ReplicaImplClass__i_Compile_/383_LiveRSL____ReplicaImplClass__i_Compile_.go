// Package _383_LiveRSL____ReplicaImplClass__i_Compile
// Dafny module _383_LiveRSL____ReplicaImplClass__i_Compile compiled into Go

package _383_LiveRSL____ReplicaImplClass__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
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
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_276_LiveRSL____ExecutorState__i_Compile "276_LiveRSL____ExecutorState__i_Compile_"
	_278_LiveRSL____LearnerState__i_Compile "278_LiveRSL____LearnerState__i_Compile_"
	_283_LiveRSL____CClockReading__i_Compile "283_LiveRSL____CClockReading__i_Compile_"
	_285_LiveRSL____ReplicaState__i_Compile "285_LiveRSL____ReplicaState__i_Compile_"
	_294_LiveRSL____MinCQuorumSize__i_Compile "294_LiveRSL____MinCQuorumSize__i_Compile_"
	_297_LiveRSL____ElectionModel__i_Compile "297_LiveRSL____ElectionModel__i_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
	_301_Impl____LiveRSL____Broadcast__i_Compile "301_Impl____LiveRSL____Broadcast__i_Compile_"
	_305_LiveRSL____ProposerLemmas__i_Compile "305_LiveRSL____ProposerLemmas__i_Compile_"
	_308_LiveRSL____ProposerModel__i_Compile "308_LiveRSL____ProposerModel__i_Compile_"
	_30_Collections____Seqs__i_Compile "30_Collections____Seqs__i_Compile_"
	_312_LiveRSL____AcceptorModel__i_Compile "312_LiveRSL____AcceptorModel__i_Compile_"
	_316_LiveRSL____LearnerModel__i_Compile "316_LiveRSL____LearnerModel__i_Compile_"
	_321_LiveRSL____ExecutorModel__i_Compile "321_LiveRSL____ExecutorModel__i_Compile_"
	_323_LiveRSL____ReplicaModel__Part1__i_Compile "323_LiveRSL____ReplicaModel__Part1__i_Compile_"
	_32_Math____mul__nonlinear__i_Compile "32_Math____mul__nonlinear__i_Compile_"
	_331_LiveRSL____ReplicaModel__Part2__i_Compile "331_LiveRSL____ReplicaModel__Part2__i_Compile_"
	_339_LiveRSL____ReplicaModel__Part3__i_Compile "339_LiveRSL____ReplicaModel__Part3__i_Compile_"
	_347_LiveRSL____ReplicaModel__Part4__i_Compile "347_LiveRSL____ReplicaModel__Part4__i_Compile_"
	_34_Math____mul__auto__proofs__i_Compile "34_Math____mul__auto__proofs__i_Compile_"
	_355_LiveRSL____ReplicaModel__Part5__i_Compile "355_LiveRSL____ReplicaModel__Part5__i_Compile_"
	_357_LiveRSL____ReplicaModel__i_Compile "357_LiveRSL____ReplicaModel__i_Compile_"
	_368_LiveRSL____UdpRSL__i_Compile "368_LiveRSL____UdpRSL__i_Compile_"
	_36_Math____mul__auto__i_Compile "36_Math____mul__auto__i_Compile_"
	_372_LiveRSL____QRelations__i_Compile "372_LiveRSL____QRelations__i_Compile_"
	_374_LiveRSL____ReplicaImplLemmas__i_Compile "374_LiveRSL____ReplicaImplLemmas__i_Compile_"
	_40_Math____mul__i_Compile "40_Math____mul__i_Compile_"
	_42_Math____div__nonlinear__i_Compile "42_Math____div__nonlinear__i_Compile_"
	_44_Math____mod__auto__proofs__i_Compile "44_Math____mod__auto__proofs__i_Compile_"
	_46_Math____mod__auto__i_Compile "46_Math____mod__auto__i_Compile_"
	_50_AppStateMachine__i_Compile "50_AppStateMachine__i_Compile_"
	_54_Concrete__NodeIdentity__i_Compile "54_Concrete__NodeIdentity__i_Compile_"
	_56_LiveRSL____Types__i_Compile "56_LiveRSL____Types__i_Compile_"
	_58_Collections____Sets__i_Compile "58_Collections____Sets__i_Compile_"
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
	_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
	_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
	_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
	_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
	_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
	_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
	_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
	_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
	_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
	_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
	_System "System_"
	"clock"
	_dafny "dafny"
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
var _ _316_LiveRSL____LearnerModel__i_Compile.Dummy__
var _ _321_LiveRSL____ExecutorModel__i_Compile.Dummy__
var _ _323_LiveRSL____ReplicaModel__Part1__i_Compile.Dummy__
var _ _331_LiveRSL____ReplicaModel__Part2__i_Compile.Dummy__
var _ _339_LiveRSL____ReplicaModel__Part3__i_Compile.Dummy__
var _ _347_LiveRSL____ReplicaModel__Part4__i_Compile.Dummy__
var _ _355_LiveRSL____ReplicaModel__Part5__i_Compile.Dummy__
var _ _357_LiveRSL____ReplicaModel__i_Compile.Dummy__
var _ _368_LiveRSL____UdpRSL__i_Compile.Dummy__
var _ _372_LiveRSL____QRelations__i_Compile.Dummy__
var _ _374_LiveRSL____ReplicaImplLemmas__i_Compile.Dummy__

type Dummy__ struct{}

// Definition of class ReplicaImpl
type ReplicaImpl struct {
	Replica               _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
	NextActionIndex       uint64
	UdpClient             *_9_Native____Io__s_Compile.UdpClient
	LocalAddr             _9_Native____Io__s_Compile.EndPoint
	Cur__req__set         *_9_Native____Io__s_Compile.MutableSet
	Prev__req__set        *_9_Native____Io__s_Compile.MutableSet
	Reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap
	Msg__grammar          _176_Common____GenericMarshalling__i_Compile.G
}

func New_ReplicaImpl_() *ReplicaImpl {
	_this := ReplicaImpl{}

	_this.Replica = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	_this.NextActionIndex = 0
	_this.UdpClient = (*_9_Native____Io__s_Compile.UdpClient)(nil)
	_this.LocalAddr = _9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint)
	_this.Cur__req__set = (*_9_Native____Io__s_Compile.MutableSet)(nil)
	_this.Prev__req__set = (*_9_Native____Io__s_Compile.MutableSet)(nil)
	_this.Reply__cache__mutable = (*_9_Native____Io__s_Compile.MutableMap)(nil)
	_this.Msg__grammar = _176_Common____GenericMarshalling__i_Compile.Type_G_().Default().(_176_Common____GenericMarshalling__i_Compile.G)
	return &_this
}

type CompanionStruct_ReplicaImpl_ struct {
}

var Companion_ReplicaImpl_ = CompanionStruct_ReplicaImpl_{}

func (_this *ReplicaImpl) Equals(other *ReplicaImpl) bool {
	return _this == other
}

func (_this *ReplicaImpl) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*ReplicaImpl)
	return ok && _this.Equals(other)
}

func (*ReplicaImpl) String() string {
	return "_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl"
}

func Type_ReplicaImpl_() _dafny.Type {
	return type_ReplicaImpl_{}
}

type type_ReplicaImpl_ struct {
}

func (_this type_ReplicaImpl_) Default() interface{} {
	return (*ReplicaImpl)(nil)
}

func (_this type_ReplicaImpl_) String() string {
	return "_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl"
}
func (_this *ReplicaImpl) Ctor__() {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _5234_empty__Udp *_9_Native____Io__s_Compile.UdpClient
	var _ = _5234_empty__Udp
	var _nw13 = _9_Native____Io__s_Compile.New_UdpClient_() // this initializes all UdpClient fields to nil
	var _ = _nw13
	_nw13.Ctor__() // this does literally nothing
	_5234_empty__Udp = _nw13
	(_this).UdpClient = _5234_empty__Udp // this.UdpClient set to an instance with nil fields
	var _5235_empty__MutableMap *_9_Native____Io__s_Compile.MutableMap
	var _ = _5235_empty__MutableMap
	var _out404 *_9_Native____Io__s_Compile.MutableMap
	var _ = _out404
	_out404 = _9_Native____Io__s_Compile.Companion_MutableMap_.EmptyMap()
	_5235_empty__MutableMap = _out404
	(_this).Reply__cache__mutable = _5235_empty__MutableMap
	var _5236_empty__MutableSet *_9_Native____Io__s_Compile.MutableSet
	var _ = _5236_empty__MutableSet
	var _out405 *_9_Native____Io__s_Compile.MutableSet
	var _ = _out405
	_out405 = _9_Native____Io__s_Compile.Companion_MutableSet_.EmptySet()
	_5236_empty__MutableSet = _out405
	(_this).Cur__req__set = _5236_empty__MutableSet
	(_this).Prev__req__set = _5236_empty__MutableSet
}
func (_this *ReplicaImpl) ConstructUdpClient(maxQueueSw *clock.Stopwatch, constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) (bool, *_9_Native____Io__s_Compile.UdpClient) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ok bool = false
	var _ = ok
	var client *_9_Native____Io__s_Compile.UdpClient = (*_9_Native____Io__s_Compile.UdpClient)(nil)
	var _ = client
	var _5237_my__ep _9_Native____Io__s_Compile.EndPoint
	var _ = _5237_my__ep
	_5237_my__ep = ((((constants).Dtor_all()).Dtor_config()).Dtor_replica__ids()).IndexUint((constants).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint)
	var _5238_ip__byte__array *_dafny.Array
	var _ = _5238_ip__byte__array
	var _nw14 = _dafny.NewArrayWithValue(0, ((_5237_my__ep).Dtor_addr()).Cardinality())
	var _ = _nw14
	_5238_ip__byte__array = _nw14
	{
	}
	_170_Common____Util__i_Compile.Companion_Default___.SeqIntoArrayOpt((_5237_my__ep).Dtor_addr(), _5238_ip__byte__array)
	var _5239_ip__endpoint *_9_Native____Io__s_Compile.IPEndPoint = (*_9_Native____Io__s_Compile.IPEndPoint)(nil)
	var _ = _5239_ip__endpoint
	var _out406 bool
	var _ = _out406
	var _out407 *_9_Native____Io__s_Compile.IPEndPoint
	var _ = _out407
	_out406, _out407 = _9_Native____Io__s_Compile.Companion_IPEndPoint_.Construct(_5238_ip__byte__array, (_5237_my__ep).Dtor_port())
	ok = _out406
	_5239_ip__endpoint = _out407
	var _nw15 = _9_Native____Io__s_Compile.New_UdpClient_() // initializes a UdpClient with nil fields
	var _ = _nw15
	_nw15.Ctor__() // this does literally nothing
	client = _nw15
	if !(ok) {
		return ok, client
	}
	var _out408 bool
	var _ = _out408
	var _out409 *_9_Native____Io__s_Compile.UdpClient
	var _ = _out409
	_out408, _out409 = _9_Native____Io__s_Compile.Companion_UdpClient_.Construct(_5239_ip__endpoint, maxQueueSw) // the final thing that actually matters. MaxQueueSw gets passed here
	ok = _out408
	client = _out409
	{
	}
	return ok, client
}
func (_this *ReplicaImpl) Replica__Init(maxQueueSw *clock.Stopwatch, constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ok bool = false
	var _ = ok
	var _out410 bool
	var _ = _out410
	var _out411 *_9_Native____Io__s_Compile.UdpClient
	var _ = _out411
	_out410, _out411 = (_this).ConstructUdpClient(maxQueueSw, constants) // this is where the UdpClient is made
	ok = _out410
	(_this).UdpClient = _out411
	if ok {
		var _out412 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
		var _ = _out412
		var _out413 *_9_Native____Io__s_Compile.MutableSet
		var _ = _out413
		var _out414 *_9_Native____Io__s_Compile.MutableSet
		var _ = _out414
		var _out415 *_9_Native____Io__s_Compile.MutableMap
		var _ = _out415
		_out412, _out413, _out414, _out415 = _323_LiveRSL____ReplicaModel__Part1__i_Compile.Companion_Default___.InitReplicaState(constants)
		(_this).Replica = _out412
		(_this).Cur__req__set = _out413
		(_this).Prev__req__set = _out414
		(_this).Reply__cache__mutable = _out415
		(_this).NextActionIndex = uint64(0)
		(_this).LocalAddr = (((((_this.Replica).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).IndexUint(((_this.Replica).Dtor_constants()).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint)
		{
		}
		(_this).Msg__grammar = _228_LiveRSL____PacketParsing__i_Compile.Companion_Default___.CMessage__grammar()
	}
	return ok
}

// End of class ReplicaImpl
