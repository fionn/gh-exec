# gh-exec

GitHub CLI extension for executing commands inheriting the GitHub authentication context.
Use it to wrap commands that will consume a `GITHUB_TOKEN`.

## Installation

### Normal

```shell
gh extension install fionn/gh-exec
```

### Development

```shell
go build
gh extension install .
```

## Usage

```
gh exec <command> [arguments]
```

For example,

```
$ gh exec pinact run --update --check
.github/workflows/release.yaml:23
-         uses: actions/checkout@v7.0.1
+         uses: actions/checkout@3d3c42e5aac5ba805825da76410c181273ba90b1 # v7.0.1
.github/workflows/release.yaml:28
-         uses: cli/gh-extension-precompile@v2.2.0
+         uses: cli/gh-extension-precompile@76961aa3bd1123d0a6fd42d0a41aca0696937c39 # v2.2.0
```
to use `pinact` while avoiding rate limits imposed on unauthenticated users.
