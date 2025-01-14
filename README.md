# GIS App Demo


## Tools

There are three parts involved, front, back an database.

## Front End

The frontend is built using *React*, with components sourced from *Material UI*.
 State management is handled through *Redux*, and communication with the server is facilitated by *Axios*.
  Session authentication is managed using *JWT* (JSON Web Tokens).

## Back End

The REST API is developed using *Golang* with the *Gin-Gonic* framework, along with various libraries — you can refer to the *go.mod* file for details. Instead of the traditional *Gorm* ORM, *Pgx* is used due to its better handling of complex SQL queries with named parameters and stored procedures."

## Database

The database is *PostgresSql* with the *PostGis* extension installed in order to work with Geospatial data.



