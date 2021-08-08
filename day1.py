# This problem was recently asked by Google.

# Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

# For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

# Bonus: Can you do this in one pass?

import unittest

def has_two(L, k):
	diffs = set()
	for e in L:
		if e in diffs:
			return True
		diffs.add(k - e)
	return False

class TestHasTwo(unittest.TestCase):
	def test_has_two(self):
		L = [10, 15, 3, 7]
		k = 17
		self.assertTrue(has_two(L, k))


if __name__ == '__main__':
	unittest.main()