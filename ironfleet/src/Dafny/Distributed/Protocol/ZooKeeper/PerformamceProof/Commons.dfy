include "../../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "../Timestamps/TimestampedType.dfy"
include "TimestampedLS.dfy"

include "../Types.dfy"
include "../DistributedSystem.dfy"
include "../ZKEnvironment.dfy"
include "../ZKDatabase.dfy"
include "../Follower.dfy"
include "../Leader.dfy"
include "../LearnerHandler.dfy"
include "Definitions.dfy"


/* This module contains some common lemmas */
module Zookeeper_Commons {
import opened Common__SeqIsUniqueDef_i
import opened ZKTimestamp
import opened ZooKeeper_Types
import opened ZooKeeper_Environment
import opened ZooKeeper_DistributedSystem
import opened EnvironmentTCP_s
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Follower
import opened ZooKeeper_Leader
import opened ZooKeeper_LearnerHandler
import opened Zookeeper_Performance_Definitions
import opened ZooKeeper_TimestampedDistributedSystem


lemma lemma_Math_Inequality(a:nat, b:nat)
    requires a <= b + 1
    ensures b-a+1 >= 0
{}

lemma lemma_Math_Addition(a:Timestamp)
    ensures forall x:Timestamp :: x <= x + a
{}


lemma {:axiom} lemma_Math_Mult_a()
    ensures forall x | x >= 0 ::
        forall a, b | a>=0 && b>= a>= 0 :: x*b >= x*a
{}

lemma lemma_Math_Mult_b()
    ensures forall x:Timestamp, n:nat ::
    x * n + x == x * (n+1)
{}


lemma {:axiom} lemma_Math_MaxOfInequalities(a:Timestamp, b:Timestamp, x:Timestamp, y:Timestamp)
    requires x <= y
    requires a <= x && b <= y
    ensures TimeMax(a, b) <= y
{}

lemma {:axiom} lemma_Math_Inequalities_Mult()
    ensures forall t:Timestamp, x:nat :: t * x >= 0
{}


lemma {:axiom} lemma_Math_Inequalities_CommonMult(x:Timestamp, a:Timestamp, b:Timestamp)
    requires a <= b
    ensures x * a <= x * b
{}


lemma lemma_Size_of_Supeset<T>(s:set<T>) 
    ensures forall subset:set<T> | s >= subset :: |s| >= |subset|
{
    forall subset:set<T> | s >= subset
    ensures |s| >= |subset|
    {
        if |s| < |subset| {
            assert |s-subset| > 0;
        }
    }
}


lemma lemma_Size_One_Sets<T>(s:set<T>, e:T) 
    requires s >= {e}
    requires |s| <= 1
    ensures s == {e}
{
    if s != {e} {
        if s == {} {
            assert ! (s >= {e});
        } else {
            assert |s| == 1;
            forall e' | e' in s 
            ensures e' == e {
                if e' != e {
                    assert e' in s && e in s;
                    assert s >= {e', e};
                    assert |{e', e}| == 2;
                    lemma_Size_of_Supeset(s);
                }
            }
        }
    }
}

}
