# Distributed Cache Database

### Drawbacks

- HTTP is not the most efficient protocol for communication and data transfer
- Strings (JSON) are not the most efficient way for data transfer
- Updates get lost if any peer servers are down
- Updates get lost if the host becomes unavailable for the peer server resolving the summary
- Snapshot-ing can cause data loss if the process is interrupted or killed (NOT DURABLE)
- Keeping all the data in memory is inefficient and dangerous
- Set and Get operations don't have any kind of Data Consistency control (READ/WRITE acknowledgements)
- The Set operation does not have an option to control replication factor (number of copies on each peer)
- There's no partitioning strategy, like how the data is paged/stored on the server for fast WRITE and READ access

```text
GOSSIP ONLY spreads information about the nodes
DATA REPLICATION is done through REPLICATION FACTOR and CONSISTENCY LEVEL (WRITE/READ CONSISTENCY LEVEL)

For a really efficient database like Cassandra
we need to implement COMPACTION and SSTables and have an INDEX and MemTable
COMPACTION (runs in the background):
Merge multiple SSTables into bigger SSTables -> update the index
SSTables:
SSTables are sorted string tables stored on Disk once the MemTable is flushed.
When MemTable is flushed -> create/update the INDEX
MemTable:
Represents In-Memory data of the database
INDEX:
Represents a map[Key:SSTable Offset] (kept on disk)
The INDEX is kept on disk because its size can be big
SUMMARY
Represents an index for the INDEX or a map[N Keys Range (Bucket) : INDEX FILE] => map[0:index1], map[100:index2], map[200:index3]
COMMIT LOG
Append only file in case the MemTable gets lost. It is destroyed after the MemTable is flushed
```
