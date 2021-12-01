#!/bin/bash -li

dafnylong Phase2Proof_helper.i.dfy | grep -iv "Warning"

dafny Phase2Proof.i.dfy | grep -iv "Warning"

dafny Phase2Proof_toplevel.i.dfy | grep -iv "Warning"