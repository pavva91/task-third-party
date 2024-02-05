// FileUploader.go MinIO example
package services

import (
	"time"

	"github.com/pavva91/task-third-party/models"
	"github.com/pavva91/task-third-party/repositories"
)

var (
	Delegation DelegationServicer = delegation{}
)

type DelegationServicer interface {
	List(year time.Time) ([]models.Delegation, error)
	// PollDelegations Fuction that runs asynchronously for polling delegations
	// Poll(periodInSeconds uint, apiEndpoint string, quitOnError bool, errorOutCh chan<- error, quitOnErrorTrueSignalInCh <-chan struct{}) error
}

type delegation struct{}

func (s delegation) List(year time.Time) ([]models.Delegation, error) {
	if year.IsZero() {
		return repositories.Delegation.List()
	}
	return repositories.Delegation.List()
}
