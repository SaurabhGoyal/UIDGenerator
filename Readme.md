# UID Generator
A globally unique-id generator using Twitter's Snowflake logic. The core ID generation logic is present at https://github.com/SaurabhGoyal/Snowflake, this repo consists of various mechanisms to build a highly available and performant cluster for the id generation. Currently it has only one mechanism - load balanced cluster.
## Load balanced cluster
  - 1 master node called load balancer.
  - 5 server nodes that generate the actual unique id. Each node registers itself to the LB when up and ready for generating IDs.
  - LB node takes all the requests from the clients, fetches ID from a random node and returns to the client.
  - Pros -
    - Clients do not need any routing, they just connect to the LB.
  - Cons -
    - Network hop via LB - reduced availability, increased latency.
    - Only one instance of LB - bottleneck of the system.
