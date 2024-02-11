load("@rules_python//python:defs.bzl", "py_binary", "py_test")

py_binary(
    name = "td_common",
    deps = [
        "@pypi//pymongo:pkg",
    ],
    main = "main.py",
    srcs = ["main.py"],
)

py_test(
    name = "test",
    srcs = ["test.py"],
    deps = [":td_common"],
)

