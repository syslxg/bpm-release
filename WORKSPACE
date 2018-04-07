http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.3/rules_go-0.10.3.tar.gz",
    sha256 = "feba3278c13cde8d67e341a837f69a029f698d7a27ddbb2a202be7a10b22142a",
)
http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.10.1/bazel-gazelle-0.10.1.tar.gz",
    sha256 = "d03625db67e9fb0905bbd206fa97e32ae9da894fe234a493e7517fd25faec914",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

local_repository(
    name = "com_github_xoebus_rules_bosh",
    path = "/home/cb/workspace/rules_bosh",
)

http_file(
    name = "runc",
    url = "https://github.com/opencontainers/runc/releases/download/v1.0.0-rc5/runc.amd64",
    sha256 = "eaa9c9518cc4b041eea83d8ef83aad0a347af913c65337abe5b94b636183a251",
    executable = True,
)
