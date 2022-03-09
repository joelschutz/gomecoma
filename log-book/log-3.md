# Day 3

Today I'm not commiting any code and it will be more clear why in a moment.
The main objective of the development at this point was to create an API
and start to storing stuff in the database. Let's say that this plan was
way more messy that I expected.

## Using "Experimental" code

As I was reading about Ent I eventually got in this article
[Auto generate REST CRUD with Ent and ogen](https://entgo.io/blog/2022/02/15/generate-rest-crud-with-ent-and-ogen).
It was perfect for my application since I could leave the trivial
method of the API to the code generator and focus only on the more
complex ones. And on top of that I'd have an auto-generated documentation complient with OpenAPI.

The problem was that this code is not really finished and stable,
some funcionalities are still buggy. Not that it is a problem, this
is the cost of being an early adopter of any technology. Just mentioning
this problems because I think it's important to say that I should know better
than to put my hope in a library before even read the documentation.

## Fixing other people's code

As the code still with problems I'm not commiting it at this stage to
the repository. I plan to identify and possible patch the library to
try to fix the issues before proceding with the main project. Much of
my time today was dedicated to that and identified 2 problems that
affected my project:

- List Fields like `Strings` or `Ints` don't seen to be supported and
crash the code generation with a generic error message
- The encoder for List Responses seens to be waiting a pointer but
it is receiving an object instead.

The second point seens to be a simple bug in the template for the
code generator. I already tried editing the generated code and it
started working properly. The first one though is more complex and
for now I don't see a good solution for it.

In both cases I'll open an Issue in the [Ogent repository](https://github.com/ariga/ogent)
and try to work with the community in solutions for those issues.
That's all for today
