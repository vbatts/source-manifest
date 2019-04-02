# source-manifest

Data structure and utility for exporting software provenance data.

## Architecture

The "simple" consideration to generate and collect information is:
- heterogeneous distro utilities (dpkg, rpm, pacman, etc.)
- language specific package utilities (python, golang, java jars and wars, etc.)
- ideally should run self-sufficient within the container, only writing the document to stdout
- may support running outside the container with a target root filesystem path

The base `srcinfo` will run a set of enabled utilities, handing up the package
and corresponding source information.

The utilities can be tailored per distro or package interface, writing the
package to stream, that the parent tool (`srcinfo`) will aggregate into a
package set.

## Additional Metadata

- The distribution version and its channel information
- The build instructions used to arrive at the final root filesystem
- Information on build context or transient files copied into or fetched during the build

## Data structure

The bill-of-materials document will look generally like:

```json
{
        "bomversion": "0.1",
        "packages": [
                {
                        "name": "acl",
                        "format": "rpm",
                        "version": "2.2.51",
                        "release": "12.el7",
                        "arch": "x86_64",
                        "source": {
                                "format": "rpm",
                                "name": "acl-2.2.51-12.el7.src.rpm",
                                "digest": "sha256:aaabbbcccddeeefff111222333...",
                                "url": "https://my.content.store.com/..."
                        }
                },
		{
			//...
		}
        ]
}
```

## Each Step of the Build

Information per step, to arrive at your current state.
In the case of a Dockerfile, or further a [multi-stage
Dockerfile](https://blog.alexellis.io/mutli-stage-docker-builds/), this
information SHOULD be collect at each step of the build process.
For example:

```Dockerfile
FROM golang as build
WORKDIR /go/src/github.com/vbatts/myapp
COPY . .
RUN go install -tags netgo github.com/vbatts/myapp

FROM busybox as prod
EXPOSE 7777
COPY --from=build /go/bin/myapp /usr/local/bin/myapp
COPY --from=build /go/src/github.com/vbatts/myapp/run.sh /usr/local/bin/run.sh
ENTRYPOINT ["/bin/sh", "/usr/local/bin/run.sh"]
```

Every Dockerfile "command" SHOULD generate this information, but at least
`FROM`, `RUN`, `COPY`, and `ADD`.

