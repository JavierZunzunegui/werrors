# go2 proprosal: universal WrapError type

This is a go2 proposal in the context of error wrapping. 
It is a concrete follow up to the more open ended [#35929 (errors: how should formatting be handled?)](https://github.com/golang/go/issues/35929).

## Context

Error wrapping support was added in [go 1.13](https://blog.golang.org/go1.13-errors).
It did not introduce some advanced features like error formatting and stack traces,
not because these were not desirable 
([Error Values — Problem Overview](https://go.googlesource.com/proposal/+/master/design/go2draft-error-values-overview.md#error-formatting), 
[Error Printing — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-printing.md))
but because no good implementation was found that delivered them. 
This puts forward an implementation that delivers on the above, 
and which can easily introduce new feautures going forward.

## Abstract

Currently, to be a wrapping error means implementing `interface {Unwrap() error}`. 

Under this proposal, it means beaing an instance of the new `WrapError` struct, which itself contains plain, non-wrapping errors.
The key advantage is new features can be added accross all wrapped error uses by adding functionality to `WrapError`,
as opposed to introducing additional interfaces that have to be inplemented by every individual error type.

## Details

### Immediate Changes

The following are to be added to [errors](https://golang.org/pkg/errors/) package
(implementations in [werrors.go](https://github.com/JavierZunzunegui/werrors/blob/master/werrors.go)): 
- `type WrapError struct`: 
effectively a linked list of errors, this is set to become the universal type for all wrapped errors.
It supports `Is`, `As` and `Unwrap` as expected.
- `func Wrap(error, error) error`: used to produce the `WrapError` error chain. 
All error wrapping is set to be produced through this method.

This is the extent of the 'immediate' changes put forward by this proposal.

### Follow ups

The more advanced features can be built on top of these changes as demonstrated in the following orthogonal extensions.
This are in place to demostrate the ease with which such features can be delivered,
but are not complete implementations of such features.
- [error formatting](https://github.com/JavierZunzunegui/werrors/tree/master/extension/format/werrors).
- [stack frames](https://github.com/JavierZunzunegui/werrors/tree/master/extension/frame/werrors).
- [error filtering](https://github.com/JavierZunzunegui/werrors/tree/master/extension/filter/werrors).
- [efficient serialisation](https://github.com/JavierZunzunegui/werrors/tree/master/extension/optimize/werrors).

Note that the above can be introduced without requiring changes on any other error type other than `WrapError`.
More generally, in a `WrapError` compliant world adding functionality to all wrapped errors requires updating `WrapError` only. 

### Limitations and Breaking Changes

Note the migration to `WrapError` is needed before the follow ups above (and theremore, the primary gains) can be delivered, 
but `WrapError` and `Wrap` can be introduced (with no additional functionality) without any breaking change. 
This allows for a gradual migration before the ultimate breaking changes.

The primary backwards-incompatible changes are:
- Removal of `interface {Unwrap() error}`: 
it was added with go1.13 to the errors package as the fundamental method that defines wrapped errors. 
Errors implementing this need to be changed to no longer have this method nor contain other errors in their fields, 
and the wrapping be provided by the new `Wrap` method.
This represents the largest breaking change in the proposal.
A significant error in this category is [os.PathError](https://golang.org/pkg/os/#PathError).
- Limited [`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf) support:
While `%w`-suffixed strings can be supported, having it elsewhere can't.
This makes `fmt.Errorf("foo, err=%w, bar", err)` unsupportable, but `fmt.Errorf("foo, bar, err: %w", err)` is. 
Even in the suffix scenario, the requiremenet for formatting imposes for a uniform suffix syntax, 
presumably the `": %w"` that was part of an [earlier errors draft](https://go.googlesource.com/proposal/+/master/design/29934-error-values.md#changes-to).
A new `fmt.WErrorf(error, string, ...interface{})` without `%w` may be altogether more suitable. 
Note that, more than an implementation detail, such standarisation is required if an error is to support formatting.
- Phasing out type assertions:
While the preferred error type checking is via `errors.Is` and `errors.As`, type assertions are still used widely.
These will not work as expected under `WrapError`, 
since the type of the error is always `WrapError` even if the last wrapped error is of the type being asserted.
