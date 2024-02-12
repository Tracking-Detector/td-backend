load("@rules_python//python:defs.bzl", "py_binary", "py_test")

py_binary(
    name = "td_common",
    srcs = ["main.py"],
    main = "main.py",
    deps = [
        "@pypi//pymongo:pkg",
    ],
)

py_test(
    name = "td_common_test",
    srcs = ["test.py"],
    main = "test.py",
    deps = [":td_common"],
)
