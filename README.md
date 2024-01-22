# Load Balancer


https://github.com/manthanguptaa/load-balancer/assets/42516515/57bf3db1-7860-4c62-aaee-f282dc0f81af


**A Lightweight, go-based Load Balancer with Active-Passive Server Management and Periodic Health Checks.**

This project implements a simple load balancer in Golang that distributes incoming requests amongst a pool of backend servers while monitoring their health. It features:

* **Round-robin load balancing:** Requests are evenly distributed across available servers.
* **Active-passive server management:** Only healthy servers receive requests. Unhealthy servers are automatically disabled until they recover.
* **Periodic health checks:** Each server undergoes regular health checks to assess its status.

**Benefits:**

* **Improved system resilience:** By distributing traffic across multiple servers, your application becomes more resistant to individual server failures.
* **Enhanced scalability:** Easily handle increased traffic by adding more backend servers to the pool.
* **Simplified server management:** Automatic health checks and active-passive server management remove the need for manual intervention.

**Getting Started:**

1. **Requirements:** Golang 1.17+
2. **Clone Repository:**
    ```bash
    $ git clone https://github.com/manthanguptaa/load-balancer.git
    ```
    
3. **Setup backend server**
    ```bash
    $ cd be
    $ go run .
    ```
4. **Setup load balancer**
    ```bash
    $ cd lb
    $ go run .
    ```
3. **Configure Backend Servers:**
    * Update the list of backend servers in the `lb/main.go` file. 
    * Each server should be represented as a `server` struct with its URL and initial `is_active` state.

**Health Check Endpoint:**

The mock backend server exposes a `/healthcheck` endpoint. You can customize this endpoint to perform specific health checks for your application requirements. 

**Disclaimer:**

This is a simple example of a load balancer and may not be suitable for production environments with high-performance requirements. Consider your specific needs and adapt the code accordingly.
