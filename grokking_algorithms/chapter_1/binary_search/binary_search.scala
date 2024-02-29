// binary search in Scala is recursive
def binary_search(nums: Array[Int], item: Int)(low: Int = 0, high: Int = nums.length - 1): Int = {
	// print(nums)
    // If element not found
	if (low > high) then return -1

    // Getting the middle element
	var mid = low + (high - low) / 2

    // If element found
	if (nums(mid) == item) then
		return mid
    // Searching in the left half
	else if (nums(mid) > item) then
		return binary_search(nums, item)(low, mid - 1)
    // Searching in the right half
	else
		return binary_search(nums, item)(mid + 1, high)
}

object BinarySearch {
    def main(args: Array[String]): Unit = {
        val nums = Array(1, 3, 5, 7, 9)
        val index = binary_search(nums, 3)(0, nums.length - 1);
        print("Index: " + index)
    }
}