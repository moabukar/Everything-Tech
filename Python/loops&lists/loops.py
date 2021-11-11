# Example 1
# loop over colours in list that are greater than 50

colours = [11, 34, 98, 43, 45, 54, 54]
colours.sort()

for colours in colours:
    if colours > 50:
        print(colours)


# Example 2
# Loop over the colours that have integer value

colours = [11, 34.1, 98.2, 43, 45.1, 54, 54]
colours.append(10)
colours.sort()

for colours in colours:
    if isinstance(colours,int):
        print(colours)