#!/bin/bash -li

dafnylong Phase2Proof_postFail_helper.i.dfy | grep -iv "Warning"

dafny Phase2Proof_postFail.i.dfy | grep -iv "Warning"

dafny Phase2Proof_toplevel.i.dfy | grep -iv "Warning"