package model

import (
	"github.com/QuanLab/gotaho-service/database/mysql"
	"log"
	"time"
)

type Job struct {
	ID              int64     `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Description     string    `json:"description,omitempty"`
	SchedulerType   int64     `json:"scheduler_type"`
	IsRepeat        string    `json:"is_repeat,omitempty"`
	IntervalSeconds int64     `json:"interval_seconds,omitempty"`
	IntervalMinutes int64     `json:"interval_minutes,omitempty"`
	Hour            int64     `json:"hours,omitempty"`
	Minutes         int64     `json:"minutes,omitempty"`
	WeekDay         int64     `json:"week_day,omitempty"`
	DayOfMonth      int64     `json:"day_of_month,omitempty"`
	CreatedDate     time.Time `json:"created_date,omitempty"`
	ModifiedDate    time.Time `json:"modified_date,omitempty"`
	Status          int64     `json:"status"`
}

type JobsList struct {
	Jobs []Job `json:"jobs,omitempty"`
	HasNext bool `json:"has_next,omitempty"`
}

func GetJobList(limit int, offset int) JobsList {
	jobs := GetJobs(limit + 1, offset)

	var jobStatus = GetServerStatus().JobStatusList
	for i, job := range jobs {
		for _, e := range jobStatus {
			if e.JobName == job.Name && e.StatusDesc == "Running" {
				jobs[i].Status = int64(1);
			}
		}
	}
	var hasNext = false
	if len(jobs) > limit {
		jobs = jobs[:len(jobs)-1]
		hasNext = true
	}
	return JobsList{Jobs:jobs, HasNext:hasNext}
}


func GetJobs(limit int, offset int) []Job {
	query := `SELECT
			A.ID_JOB,
			NAME,
			CASE WHEN DESCRIPTION IS NULL THEN '' ELSE DESCRIPTION END 
			AS DESCRIPTION,
			CREATED_DATE,
			MODIFIED_DATE,
			CASE WHEN B.VALUE_NUM IS NULL THEN 0 ELSE B.VALUE_NUM END AS SCHEDULER_TYPE,
			CASE WHEN B.C.VALUE_STR IS NULL THEN 'N' ELSE B.C.VALUE_STR END AS IS_REPEAT,
			CASE WHEN B.VALUE_NUM IS NULL THEN 0 ELSE B.VALUE_NUM END AS INTERVAL_SECONDS,
			CASE WHEN E.VALUE_NUM  IS NULL THEN 0 ELSE E.VALUE_NUM END AS INTERVAL_MINUTES,
			CASE WHEN F.VALUE_NUM  IS NULL THEN 0 ELSE F.VALUE_NUM END AS HOUR,
			CASE WHEN G.VALUE_NUM IS NULL THEN 0 ELSE G.VALUE_NUM END AS MINUTES,
			CASE WHEN H.VALUE_NUM IS NULL THEN 0 ELSE H.VALUE_NUM END AS WEEK_DAY,
			CASE WHEN I.VALUE_NUM IS NULL THEN 0 ELSE I.VALUE_NUM END AS DAY_OF_MONTH
		FROM (
			SELECT ID_JOB, NAME, DESCRIPTION, CREATED_DATE, MODIFIED_DATE FROM R_JOB
		) AS A
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'schedulerType'
		) AS B ON  A.ID_JOB = B.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_STR FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'repeat'
		) AS C ON  A.ID_JOB = C.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'intervalSeconds'
		) AS D ON  A.ID_JOB = D.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'intervalMinutes'
		) AS E ON  A.ID_JOB = E.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'hour'
		) AS F ON  A.ID_JOB = F.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'minutes'
		) AS G ON  A.ID_JOB = G.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'weekDay'
		) AS H ON  A.ID_JOB = H.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'dayOfMonth'
		) AS I ON  A.ID_JOB = I.ID_JOB
		LIMIT ? OFFSET ?
		;`

	rows, err := mysql.DB.Query(query, limit, offset)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var result = make([]Job, 0)
	for rows.Next() {
		var job Job
		err := rows.Scan(&job.ID, &job.Name, &job.Description, &job.CreatedDate, &job.ModifiedDate, &job.SchedulerType,
			&job.IsRepeat, &job.IntervalSeconds, &job.IntervalMinutes, &job.Hour, &job.Minutes, &job.WeekDay, &job.DayOfMonth)
		if err != nil {
			log.Println(err)
		}
		result = append(result, job)
	}
	return result
}

func FindJobs(name string, limit int) []Job {
	query :=
		`SELECT
			A.ID_JOB,
			NAME,
			CASE WHEN DESCRIPTION IS NULL THEN '' ELSE DESCRIPTION END 
			AS DESCRIPTION,
			CREATED_DATE,
			MODIFIED_DATE,
			CASE WHEN B.VALUE_NUM IS NULL THEN 0 ELSE B.VALUE_NUM END AS SCHEDULER_TYPE,
			CASE WHEN B.C.VALUE_STR IS NULL THEN 'N' ELSE B.C.VALUE_STR END AS IS_REPEAT,
			CASE WHEN B.VALUE_NUM IS NULL THEN 0 ELSE B.VALUE_NUM END AS INTERVAL_SECONDS,
			CASE WHEN E.VALUE_NUM  IS NULL THEN 0 ELSE E.VALUE_NUM END AS INTERVAL_MINUTES,
			CASE WHEN F.VALUE_NUM  IS NULL THEN 0 ELSE F.VALUE_NUM END AS HOUR,
			CASE WHEN G.VALUE_NUM IS NULL THEN 0 ELSE G.VALUE_NUM END AS MINUTES,
			CASE WHEN H.VALUE_NUM IS NULL THEN 0 ELSE H.VALUE_NUM END AS WEEK_DAY,
			CASE WHEN I.VALUE_NUM IS NULL THEN 0 ELSE I.VALUE_NUM END AS DAY_OF_MONTH
		FROM (
			SELECT ID_JOB, NAME, DESCRIPTION, CREATED_DATE, MODIFIED_DATE FROM R_JOB WHERE NAME LIKE ?
		) AS A
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'schedulerType'
		) AS B ON  A.ID_JOB = B.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_STR FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'repeat'
		) AS C ON  A.ID_JOB = C.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'intervalSeconds'
		) AS D ON  A.ID_JOB = D.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'intervalMinutes'
		) AS E ON  A.ID_JOB = E.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'hour'
		) AS F ON  A.ID_JOB = F.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'minutes'
		) AS G ON  A.ID_JOB = G.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'weekDay'
		) AS H ON  A.ID_JOB = H.ID_JOB
		LEFT JOIN (
			SELECT ID_JOB, VALUE_NUM FROM R_JOBENTRY_ATTRIBUTE WHERE CODE = 'dayOfMonth'
		) AS I ON  A.ID_JOB = I.ID_JOB
		LIMIT ?
		;`

	rows, err := mysql.DB.Query(query, "%" + name + "%", limit)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var result = make([]Job, 0)
	for rows.Next() {
		var job Job
		err := rows.Scan(&job.ID, &job.Name, &job.Description, &job.CreatedDate, &job.ModifiedDate, &job.SchedulerType,
			&job.IsRepeat, &job.IntervalSeconds, &job.IntervalMinutes, &job.Hour, &job.Minutes, &job.WeekDay, &job.DayOfMonth)
		if err != nil {
			log.Println(err)
		}
		result = append(result, job)
	}
	return result
}