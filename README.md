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
go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```

Create a postgres database. 
Fill in the correct values in .env, and database.yml.
Only then run <soda migrate>.

To make a binary file:

```
make dep
```

and then:

```
make run
```