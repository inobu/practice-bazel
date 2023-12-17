.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: init
init:
	bazel run @rules_go//go -- mod init github.com/inobu/practice-bazel/bzlmod

.PHONY: tidy
tidy:
	bazel run @rules_go//go -- mod tidy -v

.PHONY: build
build:
	bazel build //...
