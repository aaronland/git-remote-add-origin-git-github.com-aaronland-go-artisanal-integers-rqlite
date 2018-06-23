# go-artisanal-integers-rqlite

No, really.

This engine is not feature complete yet and still has bugs. It should be considered experimental.

```
$> ./bin/int -engine rqlite -dsn http://localhost:4001
31
```

Also, unless I am just doing it wrong when a single (following) peer is the only node left in a rqlite cluster it does not appear to know how to promote itself to be the leader. Instead it gets trapped in an endless loop like this:

```
2017/03/31 15:28:56 [WARN] raft: Election timeout reached, restarting election
2017/03/31 15:28:56 [INFO] raft: Node at 127.0.0.1:4006 [Candidate] entering Candidate state
2017/03/31 15:28:56 [ERR] raft: Failed to make RequestVote RPC to 127.0.0.1:4004: dial tcp 127.0.0.1:4004: getsockopt: connection refused
2017/03/31 15:28:56 [ERR] raft: Failed to make RequestVote RPC to 127.0.0.1:4002: dial tcp 127.0.0.1:4002: getsockopt: connection refused
2017/03/31 15:28:56 [DEBUG] raft: Votes needed: 2
2017/03/31 15:28:56 [DEBUG] raft: Vote granted from 127.0.0.1:4006. Tally: 1
2017/03/31 15:28:57 [WARN] raft: Election timeout reached, restarting election
2017/03/31 15:28:57 [INFO] raft: Node at 127.0.0.1:4006 [Candidate] entering Candidate state
2017/03/31 15:28:57 [ERR] raft: Failed to make RequestVote RPC to 127.0.0.1:4004: dial tcp 127.0.0.1:4004: getsockopt: connection refused
2017/03/31 15:28:57 [ERR] raft: Failed to make RequestVote RPC to 127.0.0.1:4002: dial tcp 127.0.0.1:4002: getsockopt: connection refused
2017/03/31 15:28:57 [DEBUG] raft: Votes needed: 2
2017/03/31 15:28:57 [DEBUG] raft: Vote granted from 127.0.0.1:4006. Tally: 1
2017/03/31 15:28:59 [WARN] raft: Election timeout reached, restarting election
2017/03/31 15:28:59 [INFO] raft: Node at 127.0.0.1:4006 [Candidate] entering Candidate state
2017/03/31 15:28:59 [ERR] raft: Failed to make RequestVote RPC to 127.0.0.1:4004: dial tcp 127.0.0.1:4004: getsockopt: connection refused
2017/03/31 15:28:59 [ERR] raft: Failed to make RequestVote RPC to 127.0.0.1:4002: dial tcp 127.0.0.1:4002: getsockopt: connection refused
2017/03/31 15:28:59 [DEBUG] raft: Votes needed: 2
2017/03/31 15:28:59 [DEBUG] raft: Vote granted from 127.0.0.1:4006. Tally: 1
```

But maybe I am just doing it wrong?

## Schema

```
CREATE TABLE integers (id INTEGER PRIMARY KEY AUTOINCREMENT, stub CHAR(1))
CREATE UNIQUE INDEX `by_stub` ON integers(stub)
```

## See also

* https://github.com/aaronland/go-artisanal-integers
* https://github.com/rqlite/rqlite
