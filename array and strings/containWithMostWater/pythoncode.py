def maxArea(height):
    # If the list is empty, return 0
    if not height:
        return 0

    left = 0
    right = len(height) - 1
    max_area = 0

    while left < right:
        # Find the shorter of the two lines
        min_height = min(height[left], height[right])
        # Calculate the width between the lines
        width = right - left
        # Calculate area
        area = min_height * width
        # Update max area if needed
        max_area = max(max_area, area)

        # Move the pointer pointing to the shorter line
        if height[left] <= height[right]:
            left += 1
        else:
            right -= 1

    return max_area

# Example usage
heights = [1, 8, 6, 2, 5, 4, 8, 3, 7]
result = maxArea(heights)
print("Maximum water container area:", result)  # Output: 49
