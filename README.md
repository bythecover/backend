# Running the Dev environment

First run `docker compose up` then open "localhost:80" on the browser

# Architecture

This project was built with the "Ports and Adapters" architecture. A "port" can be either, some behavior the core exposes, or some behavior the core expects an adapter to be able to accomplish. For example, a caller can request the core get a vote by an id and the core expects to be able to retrieve a vote by id from a vote repository. An adapter can either be something that uses the ports the core exposes, or something that satisfies the core's expectations. For example, the http handler is an adapter that will call to the core to get a vote by id given a request. all http request transformation concerns will be handled in the http handler.yi555 The PollPostgresRepository is an adapter that implements the ability go get a vote by id from a database as the core expects. Ports are included as part of the core while adapters are not. The goal is, code read in the core should strictly resemble the business requirements while the technical details should be handled in the adapters. If the core needs access to some technical detail it should do so through a port.