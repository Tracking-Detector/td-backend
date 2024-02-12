import unittest

import td_common


class ExampleTest(unittest.TestCase):
    def test_main(self):
        self.assertEqual(4, td_common.add(2,2))


if __name__ == "__main__":
    unittest.main()