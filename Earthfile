VERSION 0.6

install:
    LOCALLY
    # RUN go install github.com/go-courier/husky/cmd/husky@1.8.1
    RUN cd ./private/cli; go build -o ../../cli main.go

run:
    LOCALLY
    RUN eval $(./private/scripts/docker-compose/target/main)

ci-pull-request:
    LOCALLY
    FOR service IN $(./private/scripts/git/target/main)
        BUILD ./${service}+lint
        BUILD ./${service}+test
        BUILD ./${service}+build
    END

ci-deploy:
    LOCALLY
    ARG env=test
    ARG DOCKER_REGISTRY
    FOR service IN $(./private/scripts/git/target/main)
        BUILD ./${service}+deploy --env=$env --DOCKER_REGISTRY=$DOCKER_REGISTRY
    END

    # BUILD ./docs+deploy
    # BUILD ./frontend/libs/components+deploy

    # RUN ./private/scripts/git/target/updateCommitId