#!/bin/bash -li

# echo ""
# echo "Verifying GenericLemmas.i.dfy"
# dafny GenericLemmas.i.dfy | grep -iv "Warning"

echo ""
echo "Verifying Phase1Proof_helper0.i.dfy"
dafny Phase1Proof_helper0.i.dfy | grep -iv "Warning"

# echo ""
# echo "Verifying Phase2Proof_helper1.i.dfy"
# dafny Phase2Proof_helper1.i.dfy | grep -iv "Warning"

# echo ""
# echo "Verifying Phase2Proof_helper2.i.dfy"
# dafny Phase2Proof_helper2.i.dfy | grep -iv "Warning"

echo ""
echo "Verifying Phase1Proof.i.dfy"
dafny Phase1Proof.i.dfy | grep -iv "Warning"

echo ""
echo "Verifying Phase1Proof_toplevel.i.dfy"
dafny Phase1Proof_toplevel.i.dfy | grep -iv "Warning"
