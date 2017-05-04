package git

import (
	"github.com/weaveworks/flux/job"
	"github.com/weaveworks/flux/update"
)

type Note struct {
	JobID job.ID      `json:"JobID"`
	Spec  update.Spec `json:"Spec"`
}
