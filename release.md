#### How To Publish

- Create and push a tag such as v0.0.1:
  - git tag -a v0.0.1 -m "v0.0.1"
  - git push origin v0.0.1
- The workflow builds binaries for all targets and attaches them to the new GitHub Release
- During CI the binaries are built with version metadata injected via:
  ```sh
  go build -trimpath -ldflags "-s -w \
    -X github.com/OneideLuizSchneider/blitzctl/internal/version.Version=${VERSION} \
    -X github.com/OneideLuizSchneider/blitzctl/internal/version.GitCommit=$(git rev-parse --short HEAD) \
    -X github.com/OneideLuizSchneider/blitzctl/internal/version.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -o blitzctl .
  ```
  Replace `${VERSION}` with the tag you are releasing if running the command locally.
