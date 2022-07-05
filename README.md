# TCP-IP Layer-4 Load Balancer
#Golang V4 Proxy TCP-IP Load Balancer Go Implementation



To briefly talk about Load Balancer and its related technologies; <br/>
As the number of users increases in current technologies, there are 2 different ways. You increase the existing system or buy new servers etc.
Increasing existing resources means -Scale Out - adding new servers means - Scale Up - current way to grow.

![resim](https://user-images.githubusercontent.com/40759486/177352382-d58a9236-99f9-4585-913f-35226e9cbf00.png)


>  Upgrading the existing system means extra processor and ram. Instead, many companies have chosen to buy new server systems.

<br/><br/><br/>

As the amount of servers increases, we have new problems. We need to distribute the traffic coming over us efficiently. <br/>
We put a load balancer in front of this new server cluster and assign the work to the servers that will do the actual work.

![resim](https://user-images.githubusercontent.com/40759486/177360620-d76d993d-f57c-489f-8117-11bbefc8cfe1.png)

> The load balancer not only distributes the load, it scans our existing systems and identifies the systems that cannot be acted upon. 
 
<br> If necessary, it cannot assign processes to the problematic system.


![resim](https://user-images.githubusercontent.com/40759486/177361679-e0489af7-ba53-4cf1-9e57-5439d8c196e4.png)

Ok, How is my load balancer? <br>
The main reason for this is that #Go is a system language. It was my choice when writing #Go Load Balancer, which compiles really fast and is a good garbage collection system.

<br>Outside of Go, Python could be done with #Django or #Flask. However, as the load on this system increases, it may give us a headache. I don't want this and the best is to do something in #Go, one of the new world languages, and keep up with the change.
<br> <br>
In the project, I will set up a Layer-4 TCP Load Balancer.
![resim](https://user-images.githubusercontent.com/40759486/177363488-dcc27a15-b1b9-4fbd-8e01-a563e8805197.png)


<br>This structure will actually act as a proxy.
The purpose of this structure is to send only the connections received from the front end to the servers at the back end. <br> <br>

It will send traffic from the frontend to a randomly chosen backend. Bytes on the LB will be transmitted to the server until the connection is disconnected. Multiple incoming connections at the same time will be handled.

<br> <br> 
To try the project run python http server, nginx etc. but we using npx packages. 
```
> npx http-server -p 5001
```
