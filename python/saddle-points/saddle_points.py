def saddle_points(matrix):
    saddles = []

    if len(matrix) > 0:
        # Check if matrix is regular
        first_row_legth = len(matrix[0])
        for k in matrix:
            if len(k) != first_row_legth:
                raise ValueError('Irregular matrix')

        num_of_rows = len(matrix)
        num_of_columns = len(matrix[0])

        for r in range(num_of_rows):
            for c in range(num_of_columns):
                num_to_check = matrix[r][c]

                # Check if saddle row
                is_saddle_row = True
                for i in range(num_of_columns):
                    if num_to_check < matrix[r][i]:
                        is_saddle_row = False

                # Check if saddle column
                is_saddle_column = True
                for i in range(num_of_rows):
                    if num_to_check > matrix[i][c]:
                        is_saddle_column = False

                if is_saddle_row and is_saddle_column:
                    saddles.append((r, c))

    return set(saddles)
