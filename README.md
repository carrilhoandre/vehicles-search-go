# Vehicles seach api 
**Version 1.0.0**

---
# Getting started
## To Run application you need Docker ;). We need run elastic search, kiibana, importdataapi and Vehicles Search Api
docker-compose up

# To seed Elastic search data, run this:
GET http://localhost:3001/api/elastic/import-data

---
# Development
## If you need refresh graphql schemas:
go run github.com/99designs/gqlgen
## Dependencies
    github.com/99designs/gqlgen v0.11.3 (Graphql)
	github.com/vektah/gqlparser/v2 v2.0.1 (Graphql)
	github.com/go-chi/chi v4.1.1 (Web api server)
	github.com/olivere/elastic/v7 v7.0.15 (Elastic search connector)
---
# How use
Query example (graphql):
POST localhost:3000/query
{
  vehicles(text:"fiesta")
  {
    makeName,
    modelName
  }
}
---
## Contributors
    - André Carrilho
---
## License & copyright
© André Carrilho
