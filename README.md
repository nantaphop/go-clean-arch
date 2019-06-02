## Another Go Clean Architecture Example
Based on well known Uncle Bob's article, [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), and sample from [Bxcodec](https://github.com/bxcodec/go-clean-arch)

## How to run
*I didn't test yet but it should be something similar to this*
```
glide install
docker-compose up
```

## Explanation
Clean Architecture as I understand, it all about split application in to independence layer which must be test separately and can swapable.

Here we have `models`, `repository`, `usecase` and `delivery` as four layers.

- `models` - contain a struct that represent things in our application.
- `repository` - provide low level operations such as CRUD. In this example we have a `MongoUserRepositoty` but we also have `UserRepository` interface then someday we might change from `Mongo` to `Mysql`
- `usecase` - contain business logic about thing. For example, `CreateUser`, `ResetPassword`, `DeactivateUser` and `GetUser`.
- `delivery` - contain delivery method such as `HTTP` or `gRPC`

One last thing, if you take a look at `app.go`, it might look mess because of all dependencies setup go here but it will be easy when write a unit test.

## TODO List
- Add unit test
- Add more usecase
- Add more delivery method
- Add http authentication


