**Q1. Brief about microservice and advantages and disadvantages of using it**<br>
Using microservices we break down the larger application into smaller independent parts so that we can perform development, add features, or test each part of the application without touching the total application source code. Usually, we breakdown them based on the Business Functionalities.<br>
Advantages:<br>
Each Microservice takes care of only one specific job.<br>
Each functionality can have its own versions or updates etc.,<br>
Each microservice can have their own tech stack<br><br>
If one of the microservice is affected, it does not affect other Microservices.<br>
<br>
Disadvantages:<br>
Configuration between the microservices might be complex.<br>
It is difficult to track which service is down if it has large number of instances.<br>
Less secure as they communicate over some Network.<br>
<br>
**Q2: REST API:**
Restful API is like an interface used by 2 nodes/ computers in network to communicate securely.<br>
REST API is Representational State Transfer Application Program Interface. <br>
Used to develop network applications.<br>
Almost always HTTP protocol.<br>
Stateless – no need to store each others data<br>
Rest used HTTP to format those requests <br>
<br>
HTTP Methods.<br>
GET – Just get data from specified resource.<br>
POST – Submit data to specified resource.<br>
PUT: Update a specified Resource.<br>
DELETE: Delete a specific resource<br>
HEAD – Same as GET but does not return body<br>
OPTIONS – Returns HTTP supported methods<br>
PATCH – Update<br>
<br>
**Q3. Microservice vs Monolithic Application**<br>
In monolithic all the functionalities are implemented in a single component but in Microservices they are separated based on functionalities.<br>
Monolithic is more secure than the Microservice coz they are implemented in the same component nut microservices communicate over network.<br>
If one functionality needs to be change total application must be restarted in monolithic but in microservices particular functionality can only be changed and deployed.<br>
In microservices each functionality can have different tech stack compared to monolithic architecture.<br>
<br>
**Q4. SOAP API**<br>
SOAP – Simple object Access Protocol<br>
 Soap API has strict rules an guide lines to follow  whereas the REST API is kind of has loose guidelines.<br>
In SOAP, API calls cannot be cached where as in REST we can.<br>
SOAP WS security with SSL support but REST supports HTTPS and SSL.<br>
SOAP needs more Bandwidth and computing power but ret needs less Resources.<br>
Soap only with XML message formats but REST with palin text HTML, XML, JSON, YAML etc., message formats.<br>
<br>

**Q5. What are handlers and Endpoints**<br>
Handlers are used to get the request and return the information requested from specific Resource or between end points.<br>
End Points:<br>
Endpoints are those where URL  are  sent to get specific information.<br>

**Q6. HTTP Requests.**<br>
HTTP requests are made by clients to servers to get some required resources.<br>
HTTP request contains:<br>
1.	A request line.<br>
2.	A series of HTTP headers, or header fields.<br>
3.	A message body, if needed.<br>
