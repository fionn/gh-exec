# gh-exec

GitHub CLI extension for executing commands inheriting the GitHub authentication context.
Use it to wrap commands that will consume a `GITHUB_TOKEN`.

## Installation

```shell
gh extension install fionn/gh-exec
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
to use `pinact` while avoiding rate limits imposed on unauthenticated users, or

```
$ gh exec dependency-graphql fionn/gh-exec
fionn/gh-exec
├── main.yaml
│   ├── actions/checkout = 7.0.1
│   ├── actions/setup-go = 7.0.0
│   ├── fionn/newline-linter = 1.0.3
│   └── golangci/golangci-lint-action = 9.3.0
└── release.yaml
    ├── actions/checkout = 7.0.1
    └── cli/gh-extension-precompile = 2.2.0
```
to inspect the dependency graph using [fionn/dependency-graphql](https://github.com/fionn/dependency-graphql) without explicitly generating a token.
