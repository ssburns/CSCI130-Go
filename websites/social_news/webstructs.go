package main

import (
	"time"
	"google.golang.org/appengine/datastore"
)

type WebSubmission struct {
	Title string
	Link string
	SubmitBy string
	Thread int64
	SubmitDateTime time.Time
	SubmissionDesc string
	Score int64
}

type PageContainer struct {
	Stories []StoryListData
	BeforeLink string
	AfterLink time.Time
}

type StoryListData struct {
	Story WebSubmission
	Key *datastore.Key
}

const (
	WebSubmissionEntityName = "webSubmission"
)

