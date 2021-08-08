# This problem was asked by Uber.

# Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

# For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

# Follow-up: what if you can't use division?

import unittest

def other_products(L):
	prod = 1
	for e in L:
		prod *= e
	res = []
	for e in L:
		res.append(prod // e)
	return res

class TestOtherProducts(unittest.TestCase):
	def setUp(self):
		self.testcases = [
			{
			"input" : [1, 2, 3, 4, 5],
			"expected" : [120, 60, 40, 30, 24]
			}
		]
	def test_other_products(self):
		L = [1, 2, 3, 4, 5]
		expected = [120, 60, 40, 30, 24]

		for testcase in self.testcases:
			actual = other_products(testcase['input'])
			self.assertEqual(testcase['expected'], actual)

if __name__ == '__main__':
	unittest.main()