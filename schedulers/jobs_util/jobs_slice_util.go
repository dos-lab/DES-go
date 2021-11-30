package jobs_util

import (
	"DES-go/simulator"
	"DES-go/util"
	"sort"
)

type JobsSliceUtil struct{}

func GetJobsSliceUtil() JobsSliceUtil {
	return JobsSliceUtil{}
}

func (u JobsSliceUtil) ReorderToSRTF(gpuType simulator.GPUType, jobs []*simulator.Job) {
	SRTFSorter := util.Sorter{
		LenFunc:  func() int { return len(jobs) },
		LessFunc: func(i, j int) bool { return jobs[i].RemainingDuration(gpuType) < jobs[j].RemainingDuration(gpuType) },
		SwapFunc: func(i, j int) {
			o := jobs[i]
			jobs[i] = jobs[j]
			jobs[j] = o
		},
	}
	if sort.IsSorted(SRTFSorter) {
		return
	}
	sort.Sort(SRTFSorter)
}

func (u JobsSliceUtil) ReorderByJobName(jobs []*simulator.Job) {
	jobNameSorter := util.Sorter{
		LenFunc:  func() int { return len(jobs) },
		LessFunc: func(i, j int) bool { return jobs[i].JobName() < jobs[j].JobName() },
		SwapFunc: func(i, j int) {
			o := jobs[i]
			jobs[i] = jobs[j]
			jobs[j] = o
		},
	}
	if sort.IsSorted(jobNameSorter) {
		return
	}
	sort.Sort(jobNameSorter)
}

func (u JobsSliceUtil) InsertJobsSlice(job *simulator.Job, idx int, jobs []*simulator.Job) []*simulator.Job {
	back := append([]*simulator.Job{}, jobs[idx:]...)
	res := append(jobs[:idx], job)
	res = append(res, back...)
	return res
}

func (u JobsSliceUtil) RemoveJobsSlice(idx int, jobs []*simulator.Job) (*simulator.Job, []*simulator.Job) {
	removed := jobs[idx]
	jobs = append(jobs[:idx], jobs[idx+1:]...)
	return removed, jobs
}

func (u JobsSliceUtil) Copy(jobs []*simulator.Job) []*simulator.Job {
	c := make([]*simulator.Job, len(jobs))
	copy(c, jobs)
	return c
}