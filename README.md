# Golang Health Check
A booster demonstrating health checks and recovery using the kubernetes/openshift
`liveliness` and `readiness` probes.


# How to run
The project uses openshift s2i to build docker image. Run the following command
to build the docker image

```s2i build . registry.fedoraproject.org/f29/golang golang-health-check```

To start the web service, run

```docker run -p 8080:8080 golang-health-check```

The web service should now be accessible on `http://localhost:8080`
