/**
* Mirrors ZKDatabase.java in ZooKeeper.
* This class maintains the in memory database of zookeeper
* server states that includes the sessions, datatree and the
* committed logs. It is booted up  after reading the logs
* and snapshots from the disk.
*/

include "../../Impl/Common/SeqIsUniqueDef.i.dfy"

include "Types.dfy"


module ZooKeeper_ZKDatabase {
import opened Common__SeqIsUniqueDef_i

import opened ZooKeeper_Types

/*
    protected DataTree dataTree;
    protected ConcurrentHashMap<Long, Integer> sessionsWithTimeouts;
    protected FileTxnSnapLog snapLog;
    protected long minCommittedLog, maxCommittedLog;
    public static final int commitLogCount = 500;
    protected static int commitLogBuffer = 700;
    protected LinkedList<Proposal> committedLog = new LinkedList<Proposal>();
    protected ReentrantReadWriteLock logLock = new ReentrantReadWriteLock();
    volatile private boolean initialized = false;
*/

datatype ZKDatabase = ZKDatabase(
    initialized:bool,
    minCommittedLog:Zxid,
    maxCommittedLog:Zxid,
    commitLog:seq<Zxid>,
    isRunning:bool
)


function getLastLoggedZxid(db:ZKDatabase) : Zxid {
    if |db.commitLog| == 0 then NullZxid
    else db.commitLog[|db.commitLog| - 1]
}

function getInMemorySuffix(db:ZKDatabase) : seq<Zxid> 
    requires db.initialized
    requires isValidZKDatabase(db)
{
    if db.minCommittedLog == NullZxid then []
    else (
        var i, j :| 
            && 0 <= i <= j < |db.commitLog|
            && db.commitLog[i] == db.minCommittedLog 
            && db.commitLog[j] == db.maxCommittedLog;
        db.commitLog[i..j+1]
    )
}


predicate isInitializedZKDatabase(db:ZKDatabase) {
    db.initialized == true
}

predicate isWellOrderedLog(commitLog:seq<Zxid>) {
    forall i, j | 0 <= i < j < |commitLog| 
    :: ZxidLt(commitLog[i], commitLog[j])
}

predicate isValidZKDatabase(db:ZKDatabase) {
    db.initialized ==> (
        && SeqIsUnique(db.commitLog)
        && isWellOrderedLog(db.commitLog)

        && if db.minCommittedLog == NullZxid 
            then db.maxCommittedLog == NullZxid
            else    && db.maxCommittedLog != NullZxid
                    && db.minCommittedLog in db.commitLog
                    && db.maxCommittedLog in db.commitLog
                    && (db.minCommittedLog != db.maxCommittedLog ==> ZxidLt(db.minCommittedLog, db.maxCommittedLog))
                    && exists i, j :: (
                        && 0 <= i <= j < |db.commitLog|
                        && db.commitLog[i] == db.minCommittedLog 
                        && db.commitLog[j] == db.maxCommittedLog
                    )
    )
}

/* Valid initial state for a ZKDatabase */
predicate ZKDatabaseInit(db:ZKDatabase) {
    && db.initialized == true
    && db.isRunning == false
    && isValidZKDatabase(db)
}

/* Load the database from the disk onto memory and also add the transactions to the 
* committedlog in memory. */
// Don't actually need this. In ZK, loadDatabase is performed by QuorumPeer, not leader or follower */
// predicate loadDatabase(db:ZKDatabase, db':ZKDatabase, minCL:Zxid, maxCL:Zxid, cl:seq<Zxid>)  {
//     && db'.initialized == true
//     && db'.minCommittedLog == minCL
//     && db'.maxCommittedLog == maxCL
//     && db'.commitLog == cl
//     && isValidZKDatabase(db')
// }


/* Truncate the log to a given Zxid */
predicate truncDatabase(db:ZKDatabase, db':ZKDatabase, truncZxid:Zxid)  {
    if truncZxid !in db.commitLog 
    then db' == db
    else (
        var i :| 0 <= i < |db.commitLog| && db.commitLog[i] == truncZxid;
        && db'.commitLog == db.commitLog[..i+1]  // observe that log never truncates to empty, as it includes truncZxid
        && db'.initialized == db.initialized
        && db'.minCommittedLog == (if ZxidLt(db.minCommittedLog, truncZxid) then db.minCommittedLog else NullZxid)
        && db'.maxCommittedLog == (if ZxidLt(truncZxid, db.minCommittedLog) then NullZxid else truncZxid)
    )
}

/* Append a new ZooKeeper transaction to the log */
predicate commitToLog(db:ZKDatabase, db':ZKDatabase, txn:Zxid) {
    && db'.initialized == db.initialized
    && db'.commitLog == db.commitLog + [txn]
    && db'.maxCommittedLog == txn
    && db'.minCommittedLog == (
        if db.minCommittedLog == NullZxid then txn
        else db.minCommittedLog
    )
}

/* Take a snapshot */
predicate takeSnapshot (db:ZKDatabase, db':ZKDatabase) {
    db' == db.(minCommittedLog := NullZxid, maxCommittedLog := NullZxid)
}


/*****************************************************************************************
*                             Assorted Initial zkdb States                               *
*****************************************************************************************/

/* Specifies a valid initial, on-disk copy of a zkdb that Zookeeper servers load into mem
* We first specify the situation where we should send empty diff */
predicate InitialZkdbState_EmptyDiff(zkdbs: seq<ZKDatabase>) {
    && (forall db | db in zkdbs :: (
            && ZKDatabaseInit(db)
            && db.minCommittedLog == db.maxCommittedLog == NullZxid  // empty in-mem segment
        )
    ) && (forall db1, db2 | db1 in zkdbs && db2 in zkdbs :: (
            db1.commitLog == db1.commitLog    // commit logs are identical
        )
    )
}


}