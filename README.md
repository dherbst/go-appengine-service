# go-appengine-service
Example of using microservices on AppEngine.  There are two examples here, the first is a straightforward microservice topology in one single AppEngine deployment.   The second is a "versioned" service that separates the dev/qa/production environments in one AppEngine deployment so you can share the datastore among all three versions.

The first example of microservices is a bit easier to do and understand.

Let's say you have a service for counting, called countess (countess is the Count's girlfriend from sesame street, by the way).  You might have a default service be your application and you delegate some calls to countess when you want to know how many things happened.

## Running

    goapp serve src/default/dispatch.yaml src/default/app.yaml src/mydevservice/app.yaml
