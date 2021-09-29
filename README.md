### Erratum - key / value error custom error type.

Erratum is a custom error type for [cockroachdb](https://github.com/cockroachdb/errors) 
modelled after [exthttp](https://github.com/cockroachdb/errors/tree/master/exthttp).

```
err := fmt.Errorf("hello")
fields := erratum.Fields{"one": "oneval", "two": "twoval"}
err = erratum.WrapWithFields(err, fields)
```

Will be formatted like this:
```
  hello
  (1) fields: [one:oneval,two:twoval]
  Wraps: (2) hello
  Error types: (1) *erratum.withFields (2) *errors.errorString
```
### Wait, what about that proto file?

However, the Makefile and proto files here are currently lies, as I didn't figure out the tooling to generate all the code.

Do I know for sure I can protobuf-encode my values?

Generally, before you chose an implementation you need to ask yourself:
1. do you need the key-value pairs to be portable over the network?
     >    if no here, then you don't need to create an encode/decode pair
2. does your tech stack already define a standardized way to port payloads over the network?
     >    if no here, that's going to create a complex conversation within your project. It's not clear that protobuf is always the best choice. Then regardless of tech, you'll need to introduce some new processes and testing best practices to check all this.
3. if yes to both, then what I started to do with the proto file is a good idea
