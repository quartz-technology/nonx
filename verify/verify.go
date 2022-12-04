package verify

import (
	"errors"
	"github.com/0xpanoramix/frd-go/data"
	"github.com/sirupsen/logrus"
)

var (
	ErrBrokenCommitment = errors.New("proposed and committed payloads differ")
)

func Run(cfg *Configuration) error {
	// First, we retrieve the payloads sent to the proposers by the relays.
	proposerPayloadsDelivered, err := cfg.base.DC.GetProposerPayloadsDelivered(
		&data.GetProposerPayloadsDeliveredOptions{
			Slot: cfg.slot,
		},
	)
	if err != nil {
		logrus.WithError(err).Error("failed to retrieve distributed payload by relay")
		return err
	}

	// Then, we retrieve the block actually proposed to the network.
	proposedBlock, err := cfg.base.EC.GetPartialBeaconBellatrixBlock(cfg.slot)
	if err != nil {
		logrus.WithError(err).Error("failed to retrieve proposed payload")
		return err
	}

	proposedPayload := proposerPayloadsDelivered[0].BlockHash.String()
	committedPayload := proposedBlock.Body.ExecutionPayload.BlockHash

	// Finally, we compare the two block hashes.
	if proposedPayload != committedPayload {
		logrus.WithError(ErrBrokenCommitment).WithFields(
			logrus.Fields{
				"proposed_payload":  proposedPayload,
				"committed_payload": committedPayload,
			},
		).Error("commitment has not been respected by the proposer")
		return ErrBrokenCommitment
	}

	logrus.WithField(
		"proposer_public_key",
		proposerPayloadsDelivered[0].ProposerPubkey.String(),
	).Infoln("commitment has been respected by proposer")
	return nil
}
