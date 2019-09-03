/*
 * @lc app=leetcode.cn id=1146 lang=golang
 *
 * [1146] 快照数组
 */
type SnapshotArray struct {
	mtx       sync.RWMutex
	data      []int
	snapshots [][]int
	length    int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{length: length, data: make([]int, length, length), snapshots: make([][]int, 0, 0)}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.mtx.Lock()
	this.data[index] = val
	this.mtx.Unlock()
}

func (this *SnapshotArray) Snap() int {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	data := make([]int, len(this.data), len(this.data))
	copy(data, this.data)
	this.snapshots = append(this.snapshots, data)
	return len(this.snapshots) - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	this.mtx.RLock()
	result := this.snapshots[snap_id][index]
	this.mtx.RUnlock()
	return result
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

