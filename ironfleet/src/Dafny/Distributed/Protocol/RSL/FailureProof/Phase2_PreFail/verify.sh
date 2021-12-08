#!/bin/bash -li

echo ""
echo "Verifying GenericLemmas.i.dfy"
dafny GenericLemmas.i.dfy | grep -iv "Warning"

# echo ""
# echo "Verifying (long) Phase2Proof_helper0.i.dfy"
# dafnylong Phase2Proof_helper0.i.dfy | grep -iv "Warning"

echo ""
echo "Verifying Phase2Proof_helper1.i.dfy"
dafny Phase2Proof_helper1.i.dfy | grep -iv "Warning"

echo ""
echo "Verifying (long) Phase2Proof_helper2.i.dfy"
dafnylong Phase2Proof_helper2.i.dfy | grep -iv "Warning"

echo ""
echo "Verifying Phase2Proof.i.dfy"
dafny Phase2Proof.i.dfy | grep -iv "Warning"

echo ""
echo "Verifying Phase2Proof_toplevel.i.dfy"
dafny Phase2Proof_toplevel.i.dfy | grep -iv "Warning"
