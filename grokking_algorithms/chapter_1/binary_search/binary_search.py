def binary_search(list = [], item = 0):
	# low and high keep track of which
	# part of the list you'll search in
	low = 0
	high = len(list) - 1
	find = None # if the item don't exist, the default is nil

	# we don't have while in golang, instead we use for loop
	# While you haven¿t narrowed it down to one element
	while low <= high:
		# each time check the middle element
		mid = (low + high) // 2
		guess = list[mid]
		if guess == item: # Found the item
			find = mid
			break
		elif guess > item: # The guess was too high
			high = mid - 1
		else: # The guess was too low
			low = mid + 1
	# return the value
	return find

def main():
	nums = [1, 3, 5, 7, 9]
	print(binary_search(nums, 3)) # returns the index 1, element is in the second slot
	print(binary_search(nums, -1))   # return None, it wasn't found

main()
