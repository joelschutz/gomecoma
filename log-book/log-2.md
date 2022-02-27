# Day 2

I got pretty fair in the data modeling today. Ultimatly
decided to implement a graph database, but instead of using
a solution like [Dgraph](https://dgraph.io/) I found a method
based on SQLite.

## Why not Bun?

As I had said yesterday, I opted for using [Bun](https://bun.uptrace.dev/)
instead of an ORM to have more control over the data structure, allowing
me to have more complex relations between the entities. This solution
is good, but very tedious, I dont usually like to spend more time with
the DB configuration than the data itself.

Another thing that made me decide not to use Bun is the poor documentation
about migrations and the fault utils. In the official docs it recomends
a `starter-kit` with utilities and example of how to incorporate the library
in a project. The problem is that this project was not updated in a long time
and even some of the imports are broken. In a local copy the project works fine,
but I'd really wanted a better way to work with it.

## What is Ent?

I tried a ORM called [ent](https://entgo.io/), call it like that I think misses
the point thought. Instead of a shallow abstraction of the SQL queries, ent uses
the SQL Database as a storage medium and models your data as graph database.
This deep abstraction allows for a better way to look at complex relationships,
entities have `fields`(That define the entity attributes) and `edges`(That define
the relationship it has with another entity), the way ent will store and query the
database is not our responsability.

I'm my tests it works very well for what I wanted, even to the code got very verbose
with all the files the library generated. Looking at the schema in the db, the tables and columns also seen to make sense if you need to access then from outside the ORM.

The code for the first entities I created must be commited today. To
run see for yourself, try runing with `task`:

```bash
task run
```
