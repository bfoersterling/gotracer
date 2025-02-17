## types.Func (struct)

```
// A Func represents a declared function, concrete method, or abstract
// (interface) method. Its Type() is always a *Signature.
// An abstract method may belong to many interfaces due to embedding.
type Func struct {
	object
	hasPtrRecv_ bool  // only valid for methods that don't have a type yet; use hasPtrRecv() to read
	origin      *Func // if non-nil, the Func from which this one was instantiated
}
```

#### implements

`types.Object`

#### get receiver

Its possible to get the receiver via the `FullName()` method
```
FullName: (*main.cli_args).parse
```
as part of a string with other stuff.

You can get just the receiver with the package like this:
```
Signature().Recv().Type(): *foo.foo
```
