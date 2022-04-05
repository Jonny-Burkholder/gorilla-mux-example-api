# gorilla-mux-example-api
 An example of a simple API using Gorilla/Mux

This API uses an in-memory "database" (which is just a slice of users) and gives
the client several endpoints to reach the database. Users can be accessed via GET
request by pure path query. This will return users by ID. So format should be
localhost:8000/users/001 NOT localhost:8000/users/jeff

Queries can also be reached by URL encoding. For instance, to search by first name,
use localhost:8000/users/?first=jeff, or last name with ?last=bezos. A query of
localhost:8000/users returns all users, while users/ returns 404.

Users can also be put in the database via a POST request using JSON, and the API
will return 200OK on successful entry into database
