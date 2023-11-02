# Task 3: Create a Basic Web Server

## Steps to Run the Program

1. Run the following command in your terminal to start the web server:

   ```bash
   $ go run main.go
   ```

2. To test the /hello endpoint, use the following command in your terminal:

   ```bash
   $ curl http://localhost:8085/hello
   ```

   ```bash
   Output: <h1 style='color:blue'>Hello Gophers!</h1>
   ```

3. To test the /movie endpoint, use the following command in your terminal:

    ```bash
    $ curl http://localhost:8085/movie
    ```

    ```bash
    Output: {"name":"Mission Impossible Dead Reckoning","genre":"Action","director":"Christopher McQuarrie","rating":7.8}
    ```
    
Alternatively, you can visit any web browser and enter the following URLs:

- http://localhost:8085/hello
- http://localhost:8085/movie

A webpage should be visible with the desired output.



