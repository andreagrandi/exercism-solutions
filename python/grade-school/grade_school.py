class School(object):
    def __init__(self):
        self.school_roster = {}

    def add_student(self, name, grade):
        if self.school_roster.get(grade):
            self.school_roster[grade].append(name)
        else:
            self.school_roster[grade] = [name]

    def roster(self):
        students = []
        for k in sorted(self.school_roster):
            self.school_roster[k].sort()
            students += self.school_roster[k]
        return students

    def grade(self, grade_number):
        if self.school_roster.get(grade_number):
            self.school_roster[grade_number].sort()
            return self.school_roster[grade_number]
        else:
            return []
