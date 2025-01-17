// Copyright (C) 2021 ScyllaDB

package utils

import "time"

const (
	// syncTimeout is the maximum time the sync loop can last. In normal circumstances this is in the order
	// of seconds but there are special cases we need to account for. Like when the sync loop generates new keys
	// and signs certificates it can take several seconds. When the CI creates multiple clusters in parallel on
	// a constrained CPU, one cert can easily take over 30s.
	syncTimeout        = 2 * time.Minute
	imagePullTimeout   = 4 * time.Minute
	joinClusterTimeout = 3 * time.Minute

	// memberRolloutTimeout is the maximum amount of time it takes to start a scylla pod and become ready.
	// It includes the time to pull the images, copy the necessary files (sidecar), join the cluster and similar.
	memberRolloutTimeout = 30*time.Second + imagePullTimeout + joinClusterTimeout

	baseManagerSyncTimeout = 3 * time.Minute
	managerTaskSyncTimeout = 30 * time.Second

	nodeConfigRolloutTimeout = 3 * time.Minute
)
