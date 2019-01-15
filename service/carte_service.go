package service

import (
	"encoding/xml"
	"fmt"
	"github.com/QuanLab/gotaho-service/config"
	"github.com/QuanLab/gotaho-service/errors"
	"github.com/QuanLab/gotaho-service/model"
	"github.com/QuanLab/gotaho-service/utils"
	"strings"
)

func StartJob(job model.Job) error {
	var url = strings.Replace(config.Get().Carte.StartJobURL, "{{job}}", job.Name, 1)
	data := utils.GetData(url)
	var webResult model.JobStartResult
	err := xml.Unmarshal([]byte(data), &webResult)
	if err!= nil {
		return err
	}
	if strings.Compare(webResult.Result , "OK") != 0 {
		return errors.New(webResult.Message)
	}
	return nil
}


func StopJob(job model.Job) error {
	var url = config.Get().Carte.StopJobURL
	activeCarteJobs := GetActiveCarteJobByName(job.Name)
	// single instance id running
	if len(activeCarteJobs) == 1 {
		url = strings.Replace(url, "{{job}}", job.Name, 1)
		url = strings.Replace(url, "{{id}}", activeCarteJobs[0].InstanceId, 1)
		fmt.Println(url)
		data := utils.GetData(url) //stop job by get request
		var webResult model.JobStartResult
		err := xml.Unmarshal([]byte(data), &webResult)
		if err!= nil {
			return err
		}
		if strings.Compare(webResult.Result , "OK") != 0 {
			return errors.New(webResult.Message)
		}
	} else if len(activeCarteJobs) > 1 {
		errors.New("There are more than one instance runningm, do you want to stop all")
	} else {
		errors.New("Do not have any instance running, reload page to view result")
	}
	return nil
}

func StopAllInstance(job model.Job) error {
	var url = config.Get().Carte.StopJobURL
	activeCarteJobs := GetActiveCarteJobByName(job.Name)
	// single instance id running
	if len(activeCarteJobs) == 1 {
		url = strings.Replace(url, "{{job}}", job.Name, 1)
		url = strings.Replace(url, "{{id}}", activeCarteJobs[0].InstanceId, 1)
		fmt.Println(url)
		data := utils.GetData(url) //stop job by get request
		var webResult model.JobStartResult
		err := xml.Unmarshal([]byte(data), &webResult)
		if err!= nil {
			return err
		}
		if strings.Compare(webResult.Result , "OK") != 0 {
			return errors.New(webResult.Message)
		}
	} else if len(activeCarteJobs) > 1 {
		errors.New("There are more than one instance runningm, do you want to stop all")
	} else {
		errors.New("Do not have any instance running, reload page to view result")
	}
	return nil
}