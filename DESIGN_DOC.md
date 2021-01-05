# Design notes

This is the first application I've written in golang. I chose the gin framework because it seemed lightweight and intuitive, and the documentation is good. I also used https://github.com/vsouza/go-gin-boilerplate as a guide for golang standards around file structure and code style, since I'm new to golang and not up to speed with its norms.

# Extension considerations

## The size of the URL list could grow infinitely. How might you scale this beyond the memory capacity of the system? 

To grow the blocklist beyond the memory capacity of the application, I recommend using a database and a cache.

The database can be a simple document store, for instance dynamodb, if you stick with the AWS ecosystem. The entire blocklist should be written to the DB. 

To cut down on latency, the application should maintain an in-memory cache. It should cache both DB hits *and* misses. This is to avoid look-ups for commonly-requested hosts that are not blocked. For instance, assuming a commonly requested host like google.com is allowed, we want to return that decision quickly.

To ensure that the most-commonly requested hosts get the fastest responses, I reccommend using a least-recently-used cache eviction policy. Depending on the requirements for this service, we can start with a TTL of 5 minutes for all entries. We can also use a loading cache to re-fetch data in the background when this TTL is reached.

## Assume that the number of requests will exceed the capacity of a single system, describe how might you solve this, and how might this change if you have to distribute this workload to an additional region, such as Europe. 

Horizontal scaling is a good choice for increasing capacity for this application. We can run multiple instances of the application, all connecting to the same database. Each instance can keep its own in-memory cache. If we are concened about cache dilution among the multiple app instances, we can lift the cache out to a distributed service, for instance AWS elasticache.

If we need to deploy this service to multiple geos, I suggest using a cloud-based caching service, and a distributed database, such as dynamoDB. Updates to the block-list can then be made to any DB instance and get replicated to all DB instances with eventual consistency.

## What are some strategies you might use to update the service with new URLs? Updates may be as much as 5 thousand URLs a day with updates arriving every 10 minutes.

I would expose update capabilities in the service's API that require additional authenication. For instance, we could restrict it to an IP allowlist and mTLS. We could extend the existing API in a RESTful manner.

To add to the blocklist:
PUT /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}

To remove from the blocklist:
DELETE /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}

In the case of DELETEs, we could wait for cache entries to time out before they were truly gone, so we should tune the cache TTLs to an acceptible timeframe.

## You’re woken up at 3am, what are some of the things you’ll look for?

I would start with RED: request Rate, Error rate, and Duration. 

* Rate: Is the request rate unusally high? If so, could we be runinng out of capacity or exhausting resources?
* Error: Are we returning errors? If so, check logs or alerts for stack traces or error messages and debug from there.
* Duration: Are requests taking an unusally long amount of time or timing out? If so, could there be an external dependency that's down, such as a distributed cache or database?

Hopefully we have good telemetry in the app to begin answering these questions. After initially checking RED, I would dig into additional data to continue diagnosing, depending on what the RED shows. This could mean looking at an APM (is there a problem in our application code?), distributed traces (is there a problem between our service and a dependency?), system stats (are we running out of memory/threads/disk/etc? are we failing a liveness probe? are we running all the replicas we expect? are we crashlooping?)

In addition to investigating, I will also do a quick analysis of impact. If this is a severe incident, I'll follow whatever our company's protocol is, which could involve escapating to my back-up page and/or an incident coordinator, or updating a status site.

## Does that change anything you’ve done in the app?

I did not instrument my app with comprehensive telemetry. If I were to run this in production, I would report out RED to whatever monitoring service(s) we used and integrate with our APM.

## What are some considerations for the lifecycle of the app?

Data persistence and consistency is one concern. Because we could receive updates to the blocklist via the API at any time, deploys should correctly drain their connection pools and roll across instances, so that we can continue to service these updates.

If we stick with in-app caches, we'll have to live with cold caches on fresh instances. Alternately, we could use a distributed cache, if this is a problem.

Rolling deploys is also important for consumers of this API. Because this API is in the call stack for individuals' web requests, its deploys must have zero downtime.

We should have a CI/CD pipeline to ensure that a) we're running our tests and b) we're cutting down on the toil of human deployments. 

## You need to deploy a new version of this application. What would you do?

Many deployment concerns are already laid out in the question above. Here, I would add that we should avoid breaking changes of the API. We should ensure that we do not accidentally introduce a backwards incompatible change through automated integration testing. If we _must_ have a breaking change, then we should collaborate with our partners (ostensibly the proxy team) on API design and upgrade/rollout strategy.

For the actual deployment itself, ideally we have smoke tests, CI, CD, and extensive monitoring and observability in place to "just ship it" when the work is ready.

Realistically, working in a large company with many inter-service dependencies, we may have to partner with other teams to coordinate deploys and monitor roll-outs.
