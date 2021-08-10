# Solution

## Application

I slightly modified the application so Fiber will always listen to a fixed port (8080 in the example). This will allow mapping this port as teh container port for the next steps.

## Dockerize

Being the first time when i worked with a Go application I relied on Docker documentation to prepare the Dockerfile. I used the multistaged build as the resulting image was much smaller in this case.
<https://docs.docker.com/language/golang/build-images/>

To obtain the image run the following command after changing dir to `j2m-sre-test` so the code and the Docker file will be automatically picked up:
`docker build -t j2m-test .`

## Docker compose

In order to test the application I created a docker-compose.yaml file which spins two containers, one obtained from our application and the second one using the redis:alpine image.
The REDIS_URL is passed as an environment variable.
How to test:

- Build the Docker image as per the previous step and hanging dir to `j2m-sre-test`
- Start the docker compose:
`docker compose up -d`
- Open localhost:8080 in a browser, refresh a few times to see the incrementing of the counter.
- Stop the containers:
`docker compose down`

## Registry

I've manually built and published the image to the GitLab registry using the following commands.

```bash
$ docker login registry.gitlab.com
Username: betejb
Password: 
Login Succeeded

bogdan.betej@PC1819 MINGW64 /c/NTT/Extra/jobtome-gitlab/j2m-sre-test (feature/prepare-build)
$ docker build -t registry.gitlab.com/betejb/j2m-sre-test .
.....

$ docker push registry.gitlab.com/betejb/j2m-sre-test
Using default tag: latest
The push refers to repository [registry.gitlab.com/betejb/j2m-sre-test]
2deefee5a2c8: Pushed
208cc594e520: Pushed
07363fa84210: Pushed
latest: digest: sha256:02a3a7547ce78a3b8210c9dcf5c6094aa0f8c43efe6665e51b999e1cf64414c6 size: 949
```

## Kubernetes

The implementation is a very basic one, manifest files are in k8s folder
I created two deployments one for redis and one for our Go application (a simpler option would have been to create two containers in the same deployment, in such a case the redis would have been addressed by localhost:6379), two services, a namespace and an ingress object.

The ingress will route requests received at `http://j2m-sre-test.go.dev` (implicit port HTTP 80) to the `j2m-sre-test` service.
Our Go container will connect to redis via the dedicated service `redis-go`

We assume a secret (`regcred`) is created to allow the download of images from GitLab registry.
If the CI is also GitLab instead of the secret some implicit pipeline variables should be available. 

## CICD

The GitLab Pipeline is created via `.gitlab-ci.yml`
It covers the following stages:

- test
- build
- docker-image

  The first stage `test` runs a few standard Go tests.
  Next is the build stage where the Go code is built and an artefact is generated.
  These first two stages are insipred pretty much from the Gitlab template as as discussed i did not have prior experience with Go and GitLab pipelines.

  Finally I did some researchg and added a 3rd stage in which I've used the Docker in Docker (dind) to buiild the docker image and publish it to the Gitlab registry as the manual solution presented above is not really fit for a CICD environment.