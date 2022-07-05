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

The load balancer not only distributes the load, it scans our existing systems and identifies the systems that cannot be acted upon. If necessary, it cannot assign processes to the problematic system.


![resim](https://user-images.githubusercontent.com/40759486/177361679-e0489af7-ba53-4cf1-9e57-5439d8c196e4.png)

Ok, How is my load balancer? <br>
The main reason for this is that #Go is a system language. It was my choice when writing #Go Load Balancer, which compiles really fast and is a good garbage collection system.
