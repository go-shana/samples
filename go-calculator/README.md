# go-calculator: A sample Shana microservice

This is a sample microservice that implements a simple scientific calculator. It is intended to be used as a reference for building Shana microservices.

Some best practices are followed, such as:

- Use of `internal` package.
  - Move the `config.go` to an internal package to make it secret. In general, service config is private and should not be exposed to the outside world.
  - Implement all logic in an internal package. All public packages can be considered as "controllers" in the MVC sense. They should be as thin as possible and should not contain any business logic.
- Implement `Validate` and/or `Init` method to validate the data structure.
  - It's recommended to do validation in a `Validate(context.Context)` method and initialization in a `Init(context.Context)`. If a type implements any of these methods, Shana framework will call it automatically before calling the handler or finishing loading config.
  - In a validate method, we can use `github.com/go-shana/core/validator` and its subpackages to validate the request. We can also simply throw errors if anything goes wrong.
  - In a initialize method, we can set default values for the request. Again, if anything goes wrong, throw an error to let Shana knows.
- Return error code.
  - By default, Shana RPC returns error string to the client if there is any error. However, it's recommended to return error code instead of error string. This is because error string is not stable and may change in the future. Error code is stable and can be used to identify the error type.
  - Optionally, we can return an OK error with 0 code to indicate that the request is successful.
