package verify

import (
	"github.com/0xpanoramix/frd-go/data"
	"github.com/quartz-technology/nonx/common"
	"github.com/sirupsen/logrus"
)

func Run(cfg *Configuration) error {
	// Retrieves the payload delivered by the relay to the proposer.
	relayPayload, err := cfg.base.DC.GetProposerPayloadsDelivered(
		&data.GetProposerPayloadsDeliveredOptions{
			Slot: cfg.slot,
		},
	)
	if err != nil {
		logrus.WithError(err).Error("failed to retrieve delivered payload by relay")

		return err
	}

	// Then, we retrieve the block actually proposed to the network.
	proposedBlock, err := cfg.base.EC.GetPartialBeaconBellatrixBlock(cfg.slot)
	if err != nil {
		logrus.WithError(err).Error("failed to retrieve proposed payload")

		return err
	}

	// And finally we compare the two to check if the commitment holds.
	if !common.VerifyCommitmentForPayloadHashes(&relayPayload[0], proposedBlock.Body.ExecutionPayload.BlockHash) {
		return common.ErrBrokenCommitment
	}

	return nil
}
