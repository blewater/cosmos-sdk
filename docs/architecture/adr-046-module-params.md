# ADR 046: Module Params

## Changelog

- Sep 22, 2021: Initial Draft

## Status

DRAFT

## Abstract

This ADR describes an alternative approach to how Cosmos SDK modules use, interact,
and store their respective parameters.

## Context

Currently, in the Cosmos SDK, modules that require the use of parameters use the
`x/params` module. The `x/params` works by having modules define parameters,
typically via a simple `Params` structure, and registering that structure in
the `x/params` module via a unique `Subspace` that belongs to the respective
registering module. The registering module then has unique access to its respective
`Subspace`. Through this `Subspace`, the module can get and set its `Params`
structure.

In addition, the Cosmos SDK's `x/gov` module has direct support for changing
parameters on-chain via a `ParamChangeProposal` governance proposal type, where
stakeholders can vote on suggested parameter changes.

There are various tradeoffs to using the `x/params` module to manage individual
module parameters. Namely, managing parameters essentially comes for "free" in
that developers only need to define the `Params` struct, the `Subspace`, and the
various auxiliary functions, e.g. `ParamSetPairs`, on the `Params` type. However,
there are some notable drawbacks. These drawbacks include the fact that parameters
are serialized in state via JSON which is extremely slow. In addition, parameter
changes via `ParamChangeProposal` governance proposals are _stateless_. In other
words, it is currently not possible to have any state transitions in the
application during an attempt to change param(s).

## Decision

We will build off of the alignment of `x/gov` and `x/authz` work per
[#9810](https://github.com/cosmos/cosmos-sdk/pull/9810). Namely, module developers
will create one or more unique parameter data structures that must be serialized
to state. In addition, modules must implement `sdk.Msg` message(s) and their
respective handler(s) such that when a governance proposal passes, via the work
done in [#9810](https://github.com/cosmos/cosmos-sdk/pull/9810), the `x/gov`
module can execute those messages which would then change the respective
parameters and any other state mutations necessary.

Note, it is up to developers to decide how to structure their parameters and
the respective `sdk.Msg` messages. Consider the parameters currently defined in
`x/auth` using the `x/params` module for parameter management:

```protobuf
message Params {
  uint64 max_memo_characters       = 1;
  uint64 tx_sig_limit              = 2;
  uint64 tx_size_cost_per_byte     = 3;
  uint64 sig_verify_cost_ed25519   = 4;
  uint64 sig_verify_cost_secp256k1 = 5;
}
```

Developers can choose to either create a unique data structure for every field in
`Params` or they can create a single `Params` structure as outlined above in the
case of `x/auth`.

In the former approach, a `sdk.Msg` would need to be created for every single
field along with a handler. This can become burdensome if there are a lot of
parameter fields. In the latter case, there is only a single data structure and
thus only a single message handler, however, the message handler might have to be
more sophisticated in that it might need to understand what parameters are being
changed vs what parameters are untouched.

Once the `sdk.Msg` types and corresponding handlers are defined, proposals can
be made using the `x/gov` module that are authorized by the root `x/gov`
module's account.

Continuing to use `x/auth`, we demonstrate a more complete example:

```go
type Params struct {
	MaxMemoCharacters      uint64
	TxSigLimit             uint64
	TxSizeCostPerByte      uint64
	SigVerifyCostED25519   uint64
	SigVerifyCostSecp256k1 uint64
}

type MsgUpdateParams struct {
	MaxMemoCharacters      uint64
	TxSigLimit             uint64
	TxSizeCostPerByte      uint64
	SigVerifyCostED25519   uint64
	SigVerifyCostSecp256k1 uint64
}

type MsgUpdateParamsResponse struct {}

func (ms msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
  ctx := sdk.UnwrapSDKContext(goCtx)

  // verification logic...

  // persist params
  params := ParamsFromMsg(msg)
  ms.SaveParams(ctx, params)

  return &types.MsgUpdateParamsResponse{}, nil
}

func ParamsFromMsg(msg *types.MsgUpdateParams) Params {
  // ...
}
```

## Consequences

> This section describes the resulting context, after applying the decision. All consequences should be listed here, not just the "positive" ones. A particular decision may have positive, negative, and neutral consequences, but all of them affect the team and project in the future.

### Backwards Compatibility

The new method for working with module parameters is naturally not backwards
compatible with the existing `x/params` module. However, the `x/params` will
remain in the Cosmos SDK and will be marked as deprecated with no additional
functionality being added apart from potential bug fixes. Note, the `x/params`
module may be removed entirely in a future release.

### Positive

- Module parameters are serialized more efficiently
- Module parameters changes are able to be stateful

### Negative

- Module parameter UX becomes slightly more burdensome for developers:
    - Modules are now responsible for persisting and retrieving parameter state
    - Modules are now required to have unique message handlers to handle parameter
      changes per unique parameter data structure.

### Neutral

- Requires [#9810](https://github.com/cosmos/cosmos-sdk/pull/9810) to be reviewed
  and merged.

<!-- ## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section should contain a summary of issues to be solved in future iterations (usually referencing comments from a pull-request discussion).
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR. -->

## References

- https://github.com/cosmos/cosmos-sdk/pull/9810
- https://github.com/cosmos/cosmos-sdk/issues/9438
- https://github.com/cosmos/cosmos-sdk/discussions/9913