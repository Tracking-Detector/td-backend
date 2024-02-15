load("@rules_python//python:defs.bzl", "py_library")

py_library(
    name = "minio",
    srcs = [
        "storage_service.py",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@pypi//minio",
    ],
)
