// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package erc20_test

import (
	"math/big"

	auth "github.com/evmos/evmos/v20/precompiles/authorization"
	"github.com/evmos/evmos/v20/precompiles/erc20"
)

func (s *PrecompileTestSuite) TestIsTransaction() {
	s.SetupTest()

	// Queries
	method := s.precompile.Methods[erc20.BalanceOfMethod]
	s.Require().False(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.DecimalsMethod]
	s.Require().False(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.NameMethod]
	s.Require().False(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.SymbolMethod]
	s.Require().False(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.TotalSupplyMethod]
	s.Require().False(s.precompile.IsTransaction(&method))

	// Transactions
	method = s.precompile.Methods[auth.ApproveMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[auth.IncreaseAllowanceMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[auth.DecreaseAllowanceMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.TransferMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.TransferFromMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.MintMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.BurnMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.BurnFromMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
	method = s.precompile.Methods[erc20.TransferOwnershipMethod]
	s.Require().True(s.precompile.IsTransaction(&method))
}

func (s *PrecompileTestSuite) TestRequiredGas() {
	s.SetupTest()

	testcases := []struct {
		name     string
		malleate func() []byte
		expGas   uint64
	}{
		{
			name: erc20.BalanceOfMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.BalanceOfMethod, s.keyring.GetAddr(0))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasBalanceOf,
		},
		{
			name: erc20.DecimalsMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.DecimalsMethod)
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasDecimals,
		},
		{
			name: erc20.NameMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.NameMethod)
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasName,
		},
		{
			name: erc20.SymbolMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.SymbolMethod)
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasSymbol,
		},
		{
			name: erc20.TotalSupplyMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.TotalSupplyMethod)
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasTotalSupply,
		},
		{
			name: auth.ApproveMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(auth.ApproveMethod, s.keyring.GetAddr(0), big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasApprove,
		},
		{
			name: auth.IncreaseAllowanceMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(auth.IncreaseAllowanceMethod, s.keyring.GetAddr(0), big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasIncreaseAllowance,
		},
		{
			name: auth.DecreaseAllowanceMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(auth.DecreaseAllowanceMethod, s.keyring.GetAddr(0), big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasDecreaseAllowance,
		},
		{
			name: erc20.TransferMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.TransferMethod, s.keyring.GetAddr(0), big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasTransfer,
		},
		{
			name: erc20.TransferFromMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.TransferFromMethod, s.keyring.GetAddr(0), s.keyring.GetAddr(0), big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasTransfer,
		},
		{
			name: auth.AllowanceMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(auth.AllowanceMethod, s.keyring.GetAddr(0), s.keyring.GetAddr(0))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasAllowance,
		},
		{
			name: erc20.MintMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.MintMethod, s.keyring.GetAddr(0), big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasMint,
		},
		{
			name: erc20.BurnMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.BurnMethod, big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasTransfer,
		},
		{
			name: erc20.BurnFromMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.BurnFromMethod, s.keyring.GetAddr(0), big.NewInt(1))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasTransfer,
		},
		{
			name: erc20.TransferOwnershipMethod,
			malleate: func() []byte {
				bz, err := s.precompile.ABI.Pack(erc20.TransferOwnershipMethod, s.keyring.GetAddr(0))
				s.Require().NoError(err, "expected no error packing ABI")
				return bz
			},
			expGas: erc20.GasTransferOwnership,
		},
		{
			name: "invalid method",
			malleate: func() []byte {
				return []byte("invalid method")
			},
			expGas: 0,
		},
		{
			name: "input bytes too short",
			malleate: func() []byte {
				return []byte{0x00, 0x00, 0x00}
			},
			expGas: 0,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			input := tc.malleate()

			s.Require().Equal(tc.expGas, s.precompile.RequiredGas(input))
		})
	}
}
