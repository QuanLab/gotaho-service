package model

import (
	"encoding/xml"
	"github.com/QuanLab/gotaho-service/utils"
	"log"
	"github.com/QuanLab/gotaho-service/config"
)

type CarteJob struct {
	JobName    string `xml:"jobname,omitempty"`
	InstanceId string `xml:"id,omitempty"`
	StatusDesc string `xml:"status_desc,omitempty"`
}

type ServerStatus struct {
	JobStatusList []CarteJob `xml:"jobstatuslist>jobstatus,omitempty"`
}

type JobStartResult struct {
	Result     string `xml:"result,omitempty"`
	Message    string `xml:"message,omitempty"`
	InstanceID string `xml:"id,omitempty"`
}

func GetServerStatus() ServerStatus {
	var svStatus ServerStatus
	data := utils.GetData(config.Get().Carte.KettleStatus)
	err := xml.Unmarshal([]byte(data), &svStatus)
	if err != nil {
		log.Fatalln("GetServerStatus", err)
		svStatus.JobStatusList = make([]CarteJob, 0)
	}
	return svStatus
}