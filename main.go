package main

import (
	"DES-go/schedulers"
	"DES-go/schedulers/kmeans_scheduler"
	"DES-go/schedulers/types"
	"DES-go/simulator"
)

func main() {
	// scheduler := initDummyScheduler()
	// scheduler := initSJFScheduler()
	scheduler := initKMeansScheduler()
	simu := simulator.NewSimulator(scheduler,
		simulator.WithOptionFmtPrintLevel(simulator.ShortMsgPrint),
		simulator.WithOptionDataSourceCSVPath("/Users/purchaser/go/src/DES-go/cases/case_200_all.csv"),
		simulator.WithOptionLogEnabled(true),
		simulator.WithOptionLogPath("/Users/purchaser/go/src/DES-go/logs"),
		simulator.WithOptionGPUType2Count(map[types.GPUType]int{
			"V100": 10,
			"P100": 5,
			"T4":   10,
		}))
	simu.Start()
}

func initDummyScheduler() types.Scheduler {
	return schedulers.NewDummyScheduler()
}

func initSJFScheduler() types.Scheduler {
	return schedulers.NewSJFScheduler(false)
}

func initKMeansScheduler() types.Scheduler {
	return kmeans_scheduler.New(
		kmeans_scheduler.WithScheme(kmeans_scheduler.NewSimpleOneShotScheduleScheme(false, -1)),
		kmeans_scheduler.WithDistanceAlgo(kmeans_scheduler.NewMinCostDistanceAlgo(
			kmeans_scheduler.NewMinCostByBranchAndBoundAlgo(
				kmeans_scheduler.MinCostBranchAndBoundLCStandardPredictCost),
			kmeans_scheduler.NewSimpleAddCostSolverMaker(kmeans_scheduler.DDLCostTypeStrict, 1e20))),
	)
}