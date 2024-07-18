# tcplistener

 Simple TCP listener that opens ports as needed for testing.  
 It returns a simple message to validate that the port is indeed open.

 ```bash
./tcplistener 8080 1521 443
Listening on ports: 8080 1521 443 
 ```

You can use a tool like nc to test connectivity.

```bash
nc localhost 1521
Port is open. Goodbye!
```