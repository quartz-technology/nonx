package common

import (
	"github.com/flashbots/go-boost-utils/types"
	"github.com/sirupsen/logrus"
)

// VerifyCommitmentForPayloadHashes uses the relay and the proposed payloads to determine if the
// two are not the same and returns the according boolean.
// Note that in the current implementation, the most common issue leading to a broken commitment
// is the de-synchronization of the beacon client, as it might not have the proposed payload yet.
func VerifyCommitmentForPayloadHashes(relayPayload *types.BidTrace, proposedPayloadHash string) bool {
	// Store this in a variable as we use it in multiple places in this method.
	committedPayloadHash := relayPayload.BlockHash.String()

	// The two hashes differ, there is an issue.
	if proposedPayloadHash != committedPayloadHash {
		logrus.WithError(ErrBrokenCommitment).WithFields(
			logrus.Fields{
				"slot":                   relayPayload.Slot,
				"proposed_payload_hash":  proposedPayloadHash,
				"committed_payload_hash": committedPayloadHash,
			},
		).Error("❌ commitment has not been respected by the proposer")
		return false
	}

	// The two hashes are the same, the commitment holds.
	logrus.WithFields(
		logrus.Fields{
			"slot":                   relayPayload.Slot,
			"proposed_payload_hash":  proposedPayloadHash,
			"committed_payload_hash": committedPayloadHash,
		},
	).Infoln("✅ commitment has been respected by proposer")
	return true
}
