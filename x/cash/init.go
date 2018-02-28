package cash

import (
	"github.com/confio/weave"
	"github.com/confio/weave/errors"
)

const optKey = "cash"

// GenesisAccount is used to parse the json from genesis file
// use weave.Address, so address in hex, not base64
type GenesisAccount struct {
	Address weave.Address `json:"address"`
	Set
}

// Initializer fulfils the InitStater interface to load data from
// the genesis file
type Initializer struct{}

var _ weave.Initializer = Initializer{}

// FromGenesis will parse initial account info from genesis
// and save it to the database
func (Initializer) FromGenesis(opts weave.Options, kv weave.KVStore) error {
	accts := []GenesisAccount{}
	err := opts.ReadOptions(optKey, &accts)
	if err != nil {
		return err
	}
	for _, acct := range accts {
		// try to load up into a valid address
		if len(acct.Address) != weave.AddressLength {
			return errors.ErrUnrecognizedAddress(acct.Address)
		}
		recipient := GetOrCreateWallet(kv, NewKey(acct.Address))

		// validate the coins are proper and add them
		if err := acct.Set.Validate(); err != nil {
			return err
		}
		clean, err := acct.Set.Normalize()
		if err != nil {
			return err
		}
		recipient.Set = clean

		// save set up account
		recipient.Save()
	}
	return nil
}
