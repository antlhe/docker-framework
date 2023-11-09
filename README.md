# Docker framework
## Instructions
```
// Build base image
cd usf-repo
docker build -f Dockerfile.base -t docker-framework:latest .
cd -

// Build extended image and run tests
cd teamx-repo
docker build -t team-test-image:latest .
docker run --rm team-test-image:latest
```
