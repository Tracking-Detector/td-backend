import unittest

import main


class ExampleTest(unittest.TestCase):
    def test_main(self):
        self.assertEqual(4, main.add(2,2))


if __name__ == "__main__":
    unittest.main()