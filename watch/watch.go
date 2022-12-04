package watch

import (
	"errors"
	"github.com/0xpanoramix/frd-go/data"
	"github.com/flashbots/go-boost-utils/types"
	"github.com/quartz-technology/charon/common"
	"github.com/sirupsen/logrus"
	"time"
)

// Run executes the main watcher logic. For every new proposer payload delivered by the relay,
// it asks the beacon client the corresponding block to perform the comparison and therefore
// verify if the commitment holds.
func Run(cfg *Configuration) error {
	ticker := time.NewTicker(12 * time.Second)
	latestRelaySlot := uint64(0)

	// We run the analysis for each slot.
	for ; ; <-ticker.C {
		// Retrieves the last recorded delivered payload to a proposer by the relay.
		relayPayload, err := GetLatestPayloadDeliveredByRelay(cfg.base.DC)
		if err != nil {
			logrus.WithError(err).Error("failed to retrieve latest payload delivered by the relay")
			continue
		}

		// If this is not a mev-boost block, we skip and wait for the next slot.
		if latestRelaySlot == relayPayload.Slot {
			continue
		}
		latestRelaySlot = relayPayload.Slot

		// Retrieves the corresponding block from the beacon chain.
		proposedBlock, err := cfg.base.EC.GetPartialBeaconBellatrixBlock(relayPayload.Slot)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"slot": relayPayload.Slot,
			}).Error("failed to retrieve beacon block")
			continue
		}

		// Verifies if the commitment holds or not.
		common.VerifyCommitmentForPayloadHashes(relayPayload, proposedBlock.Body.ExecutionPayload.BlockHash)
	}
}

// GetLatestPayloadDeliveredByRelay uses the data transparency client to get the latest proposer
// payload delivered, which is used in the comparison with the proposed block to the network.
func GetLatestPayloadDeliveredByRelay(dc *data.TransparencyClient) (*types.BidTrace, error) {
	res, err := dc.GetProposerPayloadsDelivered(&data.GetProposerPayloadsDeliveredOptions{Limit: 1})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("no proposer payloads delivered")
	}

	return &res[0], nil
}
