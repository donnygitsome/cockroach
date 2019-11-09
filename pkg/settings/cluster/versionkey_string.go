// Code generated by "stringer -type=VersionKey"; DO NOT EDIT.

package cluster

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Version19_1-0]
	_ = x[VersionStart19_2-1]
	_ = x[VersionQueryTxnTimestamp-2]
	_ = x[VersionStickyBit-3]
	_ = x[VersionParallelCommits-4]
	_ = x[VersionGenerationComparable-5]
	_ = x[VersionLearnerReplicas-6]
	_ = x[VersionTopLevelForeignKeys-7]
	_ = x[VersionAtomicChangeReplicasTrigger-8]
	_ = x[VersionAtomicChangeReplicas-9]
	_ = x[VersionTableDescModificationTimeFromMVCC-10]
	_ = x[VersionPartitionedBackup-11]
	_ = x[Version19_2-12]
	_ = x[VersionStart20_1-13]
	_ = x[VersionContainsEstimatesCounter-14]
	_ = x[VersionChangeReplicasDemotion-15]
	_ = x[VersionSecondaryIndexColumnFamilies-16]
	_ = x[VersionNamespaceTableWithSchemas-17]
}

const _VersionKey_name = "Version19_1VersionStart19_2VersionQueryTxnTimestampVersionStickyBitVersionParallelCommitsVersionGenerationComparableVersionLearnerReplicasVersionTopLevelForeignKeysVersionAtomicChangeReplicasTriggerVersionAtomicChangeReplicasVersionTableDescModificationTimeFromMVCCVersionPartitionedBackupVersion19_2VersionStart20_1VersionContainsEstimatesCounterVersionChangeReplicasDemotionVersionSecondaryIndexColumnFamiliesVersionNamespaceTableWithSchemas"

var _VersionKey_index = [...]uint16{0, 11, 27, 51, 67, 89, 116, 138, 164, 198, 225, 265, 289, 300, 316, 347, 376, 411, 443}

func (i VersionKey) String() string {
	if i < 0 || i >= VersionKey(len(_VersionKey_index)-1) {
		return "VersionKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VersionKey_name[_VersionKey_index[i]:_VersionKey_index[i+1]]
}
