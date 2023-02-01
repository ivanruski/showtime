# Using cue for data validation

According to the cue [docs](https://cuelang.org/docs/usecases/validation/#validating-document-oriented-databases)

> Document-oriented databases like Mongo and many others are characterized by having flexible schema. Some of them, like Mongo, optionally allow schema definitions, often in the form of JSON schema
> CUE constraints can be used to verify document-oriented databases.

- [concepts-schema.cue](concepts-schema.cue) defines our schema
- We can define the fields every concept must have plus some optional
- If a json containing field which is not defined in the schema is given, we'll reject it
- If we want to allow unknown field we can do it - [see here](https://cuetorials.com/first-steps/validate-configuration/#schema-v1)
- More specifc concepts can "extend" the base concept definiton
- We can have type constraints as well as more specific constraints like sth being an UUID


