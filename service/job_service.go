package service

import (
	"github.com/QuanLab/gotaho-service/model"
)

func SearchJobsByName(name string, limit int) model.JobsList {
	jobs := model.FindJobs(name, limit)
	return model.JobsList{Jobs:jobs}
}

// get all active Carte Job by name
func GetActiveCarteJobByName(name string) []model.CarteJob {
	var result = make([]model.CarteJob, 0)
	jobStatusList := model.GetServerStatus().JobStatusList
	for _, job := range jobStatusList {
		if job.JobName == name &&  job.StatusDesc == "Running"{
			result = append(result, job)
		}
	}
	return result
}