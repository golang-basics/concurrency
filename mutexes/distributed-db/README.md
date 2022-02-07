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
