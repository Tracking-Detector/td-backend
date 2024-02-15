load("@rules_pkg//:pkg.bzl", "pkg_tar")
pkg_tar(
    name = "content_layer",
    srcs = [
          ":contact.json",
          ":features.json",
          ":hero.json",
          ":installation_guide.json",
          ":products.json",
        ],
    package_dir = "content",
    visibility = ["//visibility:public"],
)