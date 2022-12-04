package verify

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Flags registers the flags for the verify command.
func Flags(v *viper.Viper, f *pflag.FlagSet) {
	verifySlotFlag(v, f)
}

const (
	slotViperKey = "verify.slot"
	slotFlag     = "slot"
	slotEnv      = "CHARON_VERIFY_SLOT"
)

// verifySlotFlag defines a --slot flag and binds it to a viper key used to extract its
// corresponding value from the .yaml or .json files. Finally, it uses the same key to bind
// the environment variable.
func verifySlotFlag(v *viper.Viper, f *pflag.FlagSet) {
	// verify --slot <SLOT>
	f.Uint64(slotFlag, 0, "The slot to verify.")
	err := v.BindPFlag(slotViperKey, f.Lookup(slotFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(slotViperKey, slotEnv)
	cobra.CheckErr(err)
}

// GetSlotToVerify uses the bound viper key to get the value of the --slot flag.
func GetSlotToVerify(v *viper.Viper) uint64 {
	return v.GetUint64(slotViperKey)
}
