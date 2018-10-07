class Garden(object):

    def __init__(self, diagram, students=None):
        self.PLANT_TYPES = {
            'G': 'Grass',
            'C': 'Clover',
            'R': 'Radishes',
            'V': 'Violets'
        }

        if students is None:
            self.students = [
                'Alice', 'Bob', 'Charlie', 'David',
                'Eve', 'Fred', 'Ginny', 'Harriet',
                'Ileana', 'Joseph', 'Kincaid', 'Larry'
            ]
        else:
            students.sort()
            self.students = students

        self.garden = diagram.split('\n')

    def plants(self, student):
        student_index = self.students.index(student)
        plants = []

        plants.append(self.PLANT_TYPES[self.garden[0][student_index * 2]])
        plants.append(self.PLANT_TYPES[self.garden[0][student_index * 2 + 1]])
        plants.append(self.PLANT_TYPES[self.garden[1][student_index * 2]])
        plants.append(self.PLANT_TYPES[self.garden[1][student_index * 2 + 1]])

        return plants
