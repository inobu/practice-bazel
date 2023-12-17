.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: gazelle-update-repos
gazelle-update-repos:
	bazel run //:gazelle-update-repos

.PHONY: build
build:
	bazel build //...

.PHONY: copy-proto
copy-proto:
	cp -r ./bazel-bin/source/message_go_proto_/github.com/inobu/practice-bazel/proto/* ./proto
