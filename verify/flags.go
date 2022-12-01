package verify

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Flags(v *viper.Viper, f *pflag.FlagSet) {
	verifySlotFlag(v, f)
}

const (
	slotViperKey = "verify.slot"
	slotFlag     = "slot"
	slotEnv      = "CHARON_VERIFY_SLOT"
)

func verifySlotFlag(v *viper.Viper, f *pflag.FlagSet) {
	// verify --slot <SLOT>
	f.Uint64(slotFlag, 0, "The slot to verify.")
	err := v.BindPFlag(slotViperKey, f.Lookup(slotFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(slotViperKey, slotEnv)
	cobra.CheckErr(err)
}

func GetSlotToVerify(v *viper.Viper) uint64 {
	return v.GetUint64(slotViperKey)
}
