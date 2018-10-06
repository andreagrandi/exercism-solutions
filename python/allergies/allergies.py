import math


class Allergies(object):

    def __init__(self, score):
        self.score = score
        self.elements = []
        self.allergenes = {
            1: 'eggs',
            2: 'peanuts',
            4: 'shellfish',
            8: 'strawberries',
            16: 'tomatoes',
            32: 'chocolate',
            64: 'pollen',
            128: 'cats'
        }

        biggest_exponent_found = False
        exponent = 0

        while not biggest_exponent_found:
            power = int(math.pow(2, exponent))
            if power < self.score:
                exponent += 1
            elif power > self.score:
                exponent -= 1
                biggest_exponent_found = True
            else:
                biggest_exponent_found = True

        temp_score = self.score

        while temp_score > 0:
            power = int(math.pow(2, exponent))
            if power > 0:
                if temp_score // power > 0:
                    if self.allergenes.get(power):
                        self.elements.append(self.allergenes[power])
                    temp_score = temp_score - power
            exponent -= 1

    def is_allergic_to(self, item):
        return item in self.lst

    @property
    def lst(self):
        return self.elements
