def to_rna(dna_strand):
    nucleotides = ['G', 'C', 'T', 'A']
    nucleotides_complements = {
        'G': 'C',
        'C': 'G',
        'T': 'A',
        'A': 'U'
    }

    rna = ''

    for n in dna_strand:
        n = n.upper()
        if n not in nucleotides:
            raise ValueError('Invalid nucleotides')
        rna += nucleotides_complements[n]

    return rna
