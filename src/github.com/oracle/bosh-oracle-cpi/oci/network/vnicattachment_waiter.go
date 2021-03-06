package network

import (
	"errors"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"time"
)

type vnicAttachmentWaiter struct {
	connector client.Connector
	logger    boshlog.Logger

	attachedHandler AttachedStateHandler
}

type AttachedStateHandler func(attachmentID string, vnicID string)

type VnicAttachmentWaiter interface {
	WaitFor(attachmentID string) error
}

func NewVnicAttachmentWaiter(c client.Connector, l boshlog.Logger, h AttachedStateHandler) VnicAttachmentWaiter {
	return &vnicAttachmentWaiter{connector: c, logger: l, attachedHandler: h}
}

func (w *vnicAttachmentWaiter) WaitFor(attachmentID string) (err error) {

	getAttachmentState := func() (bool, error) {

		p := compute.NewGetVnicAttachmentParams().WithVnicAttachmentID(attachmentID)
		r, err := w.connector.CoreSevice().Compute.GetVnicAttachment(p)
		if err != nil {
			return false, err
		}
		switch *r.Payload.LifecycleState {
		case "DETACHING", "DETACHED":
			return false, fmt.Errorf("Vnic Attachment %s unexpectedly terminated", attachmentID)
		case "ATTACHED":
			if w.attachedHandler != nil {
				w.attachedHandler(attachmentID, r.Payload.VnicID)
			}
			return true, nil
		case "ATTACHING":
			return true, errors.New("Waiting")
		}
		return false, errors.New("Unknown state")
	}

	retryable := boshretry.NewRetryable(getAttachmentState)
	retryStrategy := boshretry.NewAttemptRetryStrategy(120, 1*time.Second, retryable, w.logger)

	w.logger.Debug(networkLogTag, "Waiting for attachment to reach ATTACHED state...")
	if err := retryStrategy.Try(); err != nil {
		w.logger.Debug(networkLogTag, "Error waiting to reach desired state %v. Giving up.", err)
		return err
	}
	w.logger.Debug(networkLogTag, "Reached desired state")
	return nil
}
