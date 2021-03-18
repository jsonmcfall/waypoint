# How To Use

Waypoint uses `go-changelog` to generate its changelog on release:

* https://github.com/hashicorp/go-changelog

To install, run the following command:

```
go get github.com/hashicorp/go-changelog/cmd/changelog-build
```

## How to generate CHANGELOG entries for release

Below is an example for running `go-changelog` to generate a collection of
entries. It will generate output that can be inserted into CHANGELOG.md.

For more information as to what each flag does, make sure to run `changelog-build -help`.

```
changelog-build -last-release v0.5.0 -entries-dir .changelog/ -changelog-template changelog.tmpl -note-template note.tmpl -this-release 86b6b38faa7c69f26f1d4c71e271cd4285daadf9
```

## CHANGELOG entry examples

CHANGELOG entries are expected to be txt files created inside this folder
`.changelog`. The file name is expected to be the same issue number that will
be linked when the CHANGELOG is generated. So for example, if your issue is
\#1234, your file name would be `.changelog/1234.txt`.

Below are some examples of how to generate a CHANGELOG entry with your pull
request.

### Improvement

~~~
```release-note:improvement
internal/server: Add new option for configs
```
~~~

### Feature

~~~
```release-note:feature
platform/nomad: New feature integration
```
~~~

### Bug

~~~
```release-note:bug
platform/docker: Fix broken code
```
~~~

### Multiple Entries

~~~
```release-note:bug
platform/docker: Fix broken code
```

```release-note:bug
platform/nomad: Fix broken code
```

```release-note:bug
platform/k8s: Fix broken code
```
~~~

### Long Description with Markdown

~~~
```release-note:feature
main: Lorem ipsum dolor `sit amet`, _consectetur_ adipiscing elit, **sed** do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
```
~~~