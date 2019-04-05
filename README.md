# source-manifest

Data structure and utility for exporting software provenance data.

**This is a discussion of approach. The README and structs in `./types/`**
(the `./cmd/srcinfo` is barely a stub)

## Architecture

The "simple" consideration to generate and collect information is:
- heterogeneous distro utilities (dpkg, rpm, pacman, etc.)
- language specific package utilities (python, golang, java jars and wars, etc.)
- ideally should run self-sufficient within the container, only writing the document to stdout
- may support running outside the container with a target root filesystem path

The base `srcinfo` will run a set of enabled utilities, handing up the package and corresponding source information.

The utilities can be tailored per distro or package interface, writing the package to stream, that the parent tool (`srcinfo`) will aggregate into a package set.

## Additional Metadata

- The distribution version and its channel information
- The build instructions used to arrive at the final root filesystem
- Information on build context or transient files copied into or fetched during the build

## Data structure

### Materials Data

The bill-of-materials document will look generally like:

```json
{
  "struct_type": "materials",
  "packages": [
    {
      "name": "acl",
      "format": "rpm",
      "version": "2.2.51-12.el7",
      "arch": "x86_64",
      "source": [
        {
        "format": "rpm",
        "name": "acl-2.2.51-12.el7.src.rpm",
        "digest": "sha256:aaabbbcccddeeefff111222333...",
        "url": "https://my.content.store.com/..."
        }
      ]
    },
    {
      //...
    }
  ]
}
```

### Step Data

```json
{
  "struct_type": "step",
  "uuid": "49c4ea15-4696-4635-979a-36ff82a22013",
  "time": "2019-04-03 15:09:56-04:00",
  "operation": "FROM",
  "action": "docker.io/debian:latest@sha256:9a1b6b1073bf12428a55c54e6e3bb001946afbcf49b7fea6a02d345790356998"
}
```

### Host Data

```json
{
  "struct_type": "host",
  "name": "Fedora",
  "pretty_name": "Fedora 29 (Server Edition)",
  "id": "fedora",
  "version": "29 (Server Edition)",
  "version_id": "29",
  "home_url": "https://fedoraproject.org/",
  "support_url": "https://fedoraproject.org/wiki/Communicating_and_getting_help",
  "bug_report_url": "https://bugzilla.redhat.com/",
  "annotations": {
    "cpe_name": "cpe:/o:fedoraproject:fedora:29",
    "privacy_policy_url": "https://fedoraproject.org/wiki/Legal:PrivacyPolicy"
  }
}

```

## plugins

There ought to be a directory, like `/usr/libexec/srcinfo/collectors/` that all files present that are executable are executed in name sorted order.

The executable need:
- only output JSON to stdout
- expect a single argument of the target filesystem path (i.e. `/`)
- on error or failure, write to stderr and exit non-zero

The JSON output must set a top level field of `struct_type`, so the document is aggregated correctly.
The logic will do a first pass on this `struct_type` field to detect type, then a second pass to marshal specific type (given type or versioning) and where to group it.

## Each Step of the Build

Information per step, to arrive at your current state.
In the case of a Dockerfile, or further a [multi-stage Dockerfile](https://blog.alexellis.io/mutli-stage-docker-builds/), this information SHOULD be collect at each step of the build process.
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

