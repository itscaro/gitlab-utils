## Guide
If you want to use this command without docker you can fetch the binary as below
```shell script
docker create --name gitlab-utils itscaro/gitlab-utils
docker cp gitlab-utils:/app/cli ./cli
docker rm -f gitlab-utils
```

## Labeler
This command calls Gitlab API to get changes and apply labels if it finds any matching rules in `label.yaml`

You need a Gitlab token and a configuration file mounted at `/app/label.yaml` inside the container.

### An example of rules `label.yaml`
```yaml
BACK:
  - cli/**/*
  - src/**/*

CONFIG:
  - config/**/*
```

### An example job in `.gitlab-ci.yaml`
```yaml
stages:
  - label

label:
  stage: label
  script:
    - docker run --rm
      -e GITLAB_ENDPOINT=$CI_API_V4_URL
      -e GITLAB_TOKEN=<Token to access Gitlab API>
      -v "$(pwd)/build/binary:/assets/binary"
      itscaro/gitlab-utils
      label
  only:
    - merge_requests
```

## Upload file as asset
This command is useful to release assets

You need a Gitlab token and mount the file to upload inside the container.

### An example job in `.gitlab-ci.yaml`
```yaml
stages:
  - release

label:
  stage: release
  script:
    - docker run --rm
      -e GITLAB_ENDPOINT=$CI_API_V4_URL
      -e GITLAB_TOKEN=<Token to access Gitlab API>
      -e GITLAB_PROJECT_URL=$CI_PROJECT_URL
      -v "$(pwd)/build/binary:/assets/binary"
      itscaro/gitlab-utils
      upload
        --projet-url $CI_PROJECT_URL
        --project $CI_PROJECT_PATH
        --tag $CI_COMMIT_TAG
        --file /assets/binary
  only:
    refs:
      - tags
```
