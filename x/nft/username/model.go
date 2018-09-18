package username

import (
	stderrors "errors"
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/x/nft"
)

const (
	BucketName = "usrnft"
)

type Token interface {
	nft.BaseNFT
	GetPubKeys() []PublicKey
	SetPubKeys(actor weave.Address, newKeys []PublicKey) error
}

func (u *UsernameToken) GetPubKeys() []PublicKey {
	if u.Details == nil {
		return nil
	}
	return u.Details.Keys
}
func (u *UsernameToken) SetPubKeys(actor weave.Address, newKeys []PublicKey) error {
	if !u.OwnerAddress().Equals(actor) {
		panic("Not implemented, yet")
		// TODO: handle permissions
		//if !u.Base.HasApproval(actor, nft.ActionKindUpdateDetails) {
		//	return errors.ErrUnauthorized()
		//}
	}
	u.Details = &TokenDetails{Keys: newKeys}
	return nil
}

func (u *UsernameToken) OwnerAddress() weave.Address {
	return weave.Address(u.Base.Owner)
}

func (u *UsernameToken) Transfer(newOwner weave.Address) error {
	panic("implement me")
}

func (u *UsernameToken) Validate() error {
	if err := u.Base.Validate(); err != nil {
		return err
	}
	return u.Details.Validate()
}

func (u *UsernameToken) Copy() orm.CloneableData {
	panic("implement me")
}

func (t *TokenDetails) Validate() error {
	if t == nil {
		return errors.ErrInternal("must not be nil")
	}
	for _, k := range t.Keys {
		if err := k.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (p *PublicKey) Validate() error {
	const (
		minDataLength = 2
		maxDataLength = 1000
	)
	if len(p.Data) < minDataLength || len(p.Data) > maxDataLength {
		return errors.WithCode(stderrors.New("Invalid data length"), nft.CodeInvalidDataLength)
	}
	// todo: validate algorithms
	return nil
}

func (t *TokenDetails) Copy() orm.CloneableData {
	panic("implement me")
}

// AsUsername will safely type-cast any value from Bucket
func AsUsername(obj orm.Object) (Token, error) {
	if obj == nil || obj.Value() == nil {
		return nil, nil
	}
	x, ok := obj.Value().(*UsernameToken)
	if !ok {
		return nil, nft.ErrUnsupportedTokenType()
	}
	return x, nil
}
