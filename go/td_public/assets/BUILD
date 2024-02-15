load("@rules_pkg//:pkg.bzl", "pkg_tar")

pkg_tar(
    name = "assets_layer",
    package_dir = "assets",
    visibility = ["//visibility:public"],
    deps = [
        "//go/td_public/assets/content:content_layer",
    ],
)
