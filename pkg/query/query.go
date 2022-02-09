package query

import _ "embed"

// FunctionParameters extracts parameters from a function definition and
// supports a mixture of positional parameters, default value parameters,
// typed parameters*, and typed default value parameters*.
//
// * These are not valid Starlark, but we support them to enable using Python
//   type-stub files for improved editor experience.
//go:embed parameters.scm
var FunctionParameters []byte