package node

type NamespaceConfig struct {
	Name          string          `json:"name"`
	EngType       string          `json:"eng_type"`
	SnapCount     int             `json:"snap_count"`
	SnapCatchup   int             `json:"snap_catchup"`
	RaftGroupConf RaftGroupConfig `json:"raft_group_conf"`
}

type NamespaceDynamicConf struct {
}

type RaftGroupConfig struct {
	GroupID   uint64        `json:"group_id"`
	SeedNodes []ReplicaInfo `json:"seed_nodes"`
}

type MachineConfig struct {
	// server node id
	NodeID        uint64 `json:"node_id"`
	BroadcastAddr string `json:"broadcast_addr"`
	HttpAPIPort   int    `json:"http_api_port"`
	LocalRaftAddr string `json:"local_raft_addr"`
}

type ReplicaInfo struct {
	NodeID    uint64 `json:"node_id"`
	ReplicaID uint64 `json:"replica_id"`
	RaftAddr  string `json:"raft_addr"`
}

type RaftConfig struct {
	GroupID   uint64 `json:"group_id"`
	GroupName string `json:"group_name"`
	// this is replica id
	ID uint64 `json:"id"`
	// local server transport address, it
	// can be used by several raft group
	RaftAddr    string                 `json:"raft_addr"`
	DataDir     string                 `json:"data_dir"`
	WALDir      string                 `json:"wal_dir"`
	SnapDir     string                 `json:"snap_dir"`
	RaftPeers   map[uint64]ReplicaInfo `json:"raft_peers"`
	SnapCount   int                    `json:"snap_count"`
	SnapCatchup int                    `json:"snap_catchup"`
	nodeConfig  *MachineConfig
}