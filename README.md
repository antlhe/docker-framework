# Docker framework
## Instructions
```
cd docker-framework-base
docker build -f Dockerfile.base -t docker-framework:latest .
cd -
cd team_repo
docker build -t team-test-image:latest .
docker run --rm team-test-image:latest
```
