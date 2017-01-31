# octopz

A WIP for a small, simple reverse proxy.

### Goals:

- Ease of use
- Flexibility
- Speed

To achieve these goals, many features will be sacrificed.

### Main focuses:

- Dynamic updates of route tables. No restart/redeploy needed.
- Resilience
- Logging
- Simple mertric collection
- Authentication
- Rate limiting


The idea behind is to fit this into a service system, which handles loadbalancing by itself, for example put an ELB in front of the service and route using DNS. 

No loadbalancing will be provided by this service.

No SSL will be provided either, Let the load balancer in front of this service handle SSL for you.

### Concerns:

Ease of use dictates easy of deployment, however rate limiting, metric collection, etc. demands some form of storage, same is true for dynamic routing. It will be too expensive to look up in a database for routes constantly, so some form of cache needs to be implemented. Maybe an in memory tree could be used.

