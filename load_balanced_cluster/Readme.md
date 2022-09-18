# Load balanced cluster
  - 1 master node called load balancer.
  - 5 server nodes that generate the actual unique id. Each node registers itself to the LB when up and ready for generating IDs.
  - LB node takes all the requests from the clients, fetches ID from a random node and returns to the client.
  - Pros -
    - Clients do not need any routing, they just connect to the LB.
  - Cons -
    - Network hop via LB - reduced availability, increased latency.
    - Only one instance of LB - bottleneck of the system.

## Setup and run
- Install docker, docker-compose and golang.
- CLone this repo.
- Run `docker-compose up` in this directory.
- Make calls for ID generation -
```
curl localhost:8700/id
778648265552825344
```