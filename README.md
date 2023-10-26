# Goperson

Goperson is a service, which receives full name via API, enriches
answer with the most likely age, gender and nationality and saves the data in
Postgres DataBase. Upon request, provides information about found people.

- Built in Go version 1.21.3
  
## Dependencies:

- [chi]("github.com/go-chi/chi/v5")
- [godotenv](https://github.com/joho/godotenv)
- [soda](https://github.com/gobuffalo/pop)


In order to build and run this application, it is necessary to 
install Soda:

```
go install github.com/gobuffalo/pop/... 
```

create a postgres database and fill in the correct values in .env, 
and then run <soda migrate>.

To make a binary file:

```
make dep
```

and then:

```
make run
```