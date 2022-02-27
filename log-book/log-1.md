# Day 1

The first thing that I tried to do was to define the data structure
for the application. In that respect I have considerations about the data
itself and the software side of things. Let's start with that.

## Software

The question "SQL or Non-SQL" was very easy to answer. Since the kind of
data that I'm working with is highly relational(Different artists are envolved with different projects in different ways) and I needed to easily
filter and group those entities I opted for a SQL database. This choice
also helps me with storing all this data since I can use a simple SQLite
file to store all the data.

Now I needed some kind of tool to interact with the DB. My first intention
was to use [GORM](https://gorm.io/) as ORM to deal with all the queries and
migrations, but decided not to use it at the end. As I said, my data is
related in complex ways. An Artist can be a Actor in a Movie and a Producer in another, the model definition in GORM don't seen to be that
flexible. Or at least I'd have an overcomplexed model structure only that
would be a pain to extend later.

So I opted for [Bun](https://bun.uptrace.dev/), it's not that popular, but
I found it a great tool to have a more low level control over the DB. One
great thing is that I'm responsable for creating the DB migrations, this
unsures that the schema is exactly the way I wanted and help me create more complex relationships.

I admit that this aproach may not me sufficient, the way I want to
structure the data really calls for a Graph Database. If I'm able to find
a lighweight solution based on, or even similar to, SQLite, that'd be the way.

## Data Structure

The traditional way you structure a relational db is using rigid models of
data to represent entities and use foreign keys to connect those entities.
This is great, but really don't want to duplicate that in my database in
the event of an actor directing a movie or a musician acting in a TV Show.
I want to have an Artist entity that represents a person and use it as a
reference in all the posible ways it can be related to a work of media.

I could use the Artist Id as a foreign key in the Movie entity, for example,
and create a list of movies in the Artist table with the foreign key of
all the movies. This would work great in locating the artist from the Movie,
but not the other way around. If an Artist had written and directed a Movie,
they would appear twice in the `artist_movie` table, but I'd still need
to scan the movie field to find what role they had in the production.

On top of that, I really want to have a way to catalog all the artists work
in a cohesive manner so it'd be easier to see their work in different fields.
Take for example [Charlotte_Gainsbourg](https://en.wikipedia.org/wiki/Charlotte_Gainsbourg),
she has long career as an actress and a musician, it'd be great if I could
catalogue all those works and query this that in a way to easily show it to
the user.

The idea tomorow is to dive deeper in this data modeling until I find a good solution.
