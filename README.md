This command calls Gitlab API to get changes and apply labels if it finds any matching rules in `label.yml`

You need a Gitlab token and a configuration file mounted at `/app/label.yml` inside the container.

If you want to use this command without docker you can fetch the binary as below
```shell script
docker create --name gitlab-labeler itscaro/gitlab-labeler
docker cp gitlab-labeler:/app/cli ./cli
docker rm -f gitlab-labeler
```

### An example of rules `label.yml`
```yaml
BACK:
  - cli/**/*
  - src/**/*

CONFIG:
  - config/**/*
```

### An example job in `.gitlab-ci.yml`
```yaml
stages:
  - label

label:
  stage: label
  script:
    - docker run --rm -e GITLAB_ENDPOINT=$CI_API_V4_URL -e GITLAB_TOKEN=<Token to access Gitlab API> -v "$(pwd)/label.yml:/app/label.yml" itscaro/gitlab-labeler -p $CI_PROJECT_PATH -i $CI_MERGE_REQUEST_IID
  only:
    - merge_requests
```