# octopz

A WIP for a small, simple reverse proxy.  

### Goals:

- Ease of use
- Flexibility
- Speed

To achieve these goals, many features will be sacrificed.

### Main focuses:

- Dynamic updates of route tables. No restart/redeploy needed.
- Authentication
- Rate limiting
- Auto registration of endpoint when your service starts up
- Life sign checking. 

When a service starts up, it can automatically register it's schema to the proxy. 
These endpoints could be either live forever or have a TTL. 
To make sure live endpoints are not cleaned up we need life sign checking.

### Concerns:

Ease of use dictates easy of deployment, however rate limiting, metric collection, etc. demands some form of storage, same is true for dynamic routing. It will be too expensive to look up in a database for routes constantly, so some form of cache needs to be implemented. Maybe an in memory tree could be used.

Security is gonna be a concern. How do we secure the admin interface for the proxy? 
Possibility is to run this on a seperate port, and not expose this to the internet, however I do not feel this is enough. A wrongly configured proxy could lead to disasters, however implementing an authentication layer seems unfruitful...
