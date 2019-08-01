# ForgeRock SaaS Software Engineer Coding Challenge
## How to Run
exec: `$> ./run-challenge.sh`  
The project is api driven so, curl or make a request to **localhost at port 3000**.
After executing the *run-challenge.sh* there will be more instructions for interacting with the app.

### testing
- **unit tests**  
exec: `$> go test`  

- **benchmark**  
exec: `$> go test -bench .`

### API
GET, POST or DELETE at `/user/:username`

### prereqs
- go `1.12.6`
- docker `18.09.2`

*coded on a darwin/amd64 box*

## Thoughts
There is much more we could do in the way of logging and error handling. App/redis
configuration and packaging leaves something to be desired (*but perhaps we
could discuss methods of operationalization*). Durability is not here, what happens
if I lose connection to redis? How do I know that I consumed published content?

This code could easily be broken apart to by operations interacting with redis.
Publishing and subscription methods could be handled by different programs/applications.
