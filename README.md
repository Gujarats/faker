# faker [![Build Status](https://secure.travis-ci.org/Gujarats/faker.png)](http://travis-ci.org/Gujarats/faker) [![GoReportCard](https://goreportcard.com/badge/github.com/Gujarats/faker)](https://goreportcard.com/report/github.com/Gujarats/faker)
Heavily inspired by [faker](https://github.com/fzaninotto/Faker) open source project.
All data from provider is comes that project. `100%` written in Go.

## usage

```shell
// get 10 random name
randomNames := faker.Names(10)

// get a random name
randomName := faker.Name()

// get a random name with speficic format
randomName = faker.Namef("first-name-male")

// get 10 random name with specific format
randomNames = faker.Namesf(10, "first-name-male")
```

## TODO 
There are some a lot of thing left to do here like : 

 - Dummy Address
 - Dummy Phone Number
 - Dummy Company Name
