# Running the Dev environment

First run `docker compose up` then open "localhost:80" on the browser

# Architecture

This project was built with the "Ports and Adapters" architecture. A "port" can be either, some behavior the core exposes, or some behavior the core expects an adapter to be able to accomplish. For example, a caller can request the core get a vote by an id and the core expects to be able to retrieve a vote by id from a vote repository. An adapter can either be something that uses the ports the core exposes, or something that satisfies the core's expectations. For example, the http handler is an adapter that will call to the core to get a vote by id given a request. all http request transformation concerns will be handled in the http handler.yi555 The PollPostgresRepository is an adapter that implements the ability go get a vote by id from a database as the core expects. Ports are included as part of the core while adapters are not. The goal is, code read in the core should strictly resemble the business requirements while the technical details should be handled in the adapters. If the core needs access to some technical detail it should do so through a port.

# Folders

- internal/
    - adapters/ - the different ways the outside world can interact with the app, or the app can interact with the outside world.
        - http/ - allows http to be an "actor" to our app
        - persistence/ - allows our app to communicate with a database
    - core/ - core business logic
        - domain/ - Business Logic Objects
        - ports/ - interfaces that define how the service may act upon the outside world or may be acted upon
        - services/ - the implementation of business logic
        - templates/ - templ templates
            - components/ - reusable standalone components
            - pages/ - templ files meant to be used once
    - static/ - static files and hosting for the web
        - assets/ - css/image/js files

# Steps to Run Dev Environment
1. Install the latest version of Docker and Docker compose
2. Change directory to backend/
3. run `docker compose up`
