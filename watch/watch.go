package watch

import (
	"errors"
	"github.com/0xpanoramix/frd-go/data"
	"github.com/flashbots/go-boost-utils/types"
	"github.com/quartz-technology/charon/verify"
	"github.com/sirupsen/logrus"
	"time"
)

func Run(cfg *Configuration) error {
	ticker := time.NewTicker(12 * time.Second)
	latestRelaySlot := uint64(0)

	for ; ; <-ticker.C {
		relayPayload, err := GetLatestPayloadDeliveredToRelay(cfg.dc)
		if err != nil {
			return err
		}

		if latestRelaySlot == relayPayload.Slot {
			continue
		}
		latestRelaySlot = relayPayload.Slot

		proposedBlock, err := cfg.ec.GetPartialBeaconBellatrixBlock(relayPayload.Slot)
		if err != nil {
			return err
		}

		proposedPayloadHash := relayPayload.BlockHash.String()
		committedPayloadHash := proposedBlock.Body.ExecutionPayload.BlockHash

		// Finally, we compare the two block hashes.
		if proposedPayloadHash != committedPayloadHash {
			logrus.WithError(verify.ErrBrokenCommitment).WithFields(
				logrus.Fields{
					"proposed_payload_hash":  proposedPayloadHash,
					"committed_payload_hash": committedPayloadHash,
					"slot":                   relayPayload.Slot,
				},
			).Error("❌ commitment has not been respected by the proposer")
			continue
		}

		logrus.WithFields(
			logrus.Fields{
				"proposed_payload_hash":  proposedPayloadHash,
				"committed_payload_hash": committedPayloadHash,
				"slot":                   relayPayload.Slot,
			},
		).Infoln("✅ commitment has been respected by proposer")
	}
}

func GetLatestPayloadDeliveredToRelay(dc *data.TransparencyClient) (*types.BidTrace, error) {
	res, err := dc.GetProposerPayloadsDelivered(&data.GetProposerPayloadsDeliveredOptions{Limit: 1})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("no proposer payloads delivered")
	}

	return &res[0], nil
}
