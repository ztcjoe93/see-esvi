import csv
import random

# simple script to generate random samples

MIN_ROWS_PER_FILE = 3250
MAX_ROWS_PER_FILE = 4520

list_of_random_values = ['apple', 'banana', 'cataloupe', 'dragonfruit',
    'rambutan', 'durian', 'jackfruit', 'pomelo',
    'cat', 'dog', 'bird', 'dragonfly',
    'elephant', 'snake', 'ox', 'cow',
    'fire', 'water', 'earth', 'wind'
]

def generate_samples(num_of_samples: int) -> None:
    count = 1
    magnitude = num_of_samples
    while (magnitude >= 10):
        magnitude /= 10
        count += 1

    for i in range(1, num_of_samples+1):
        with open("sample_{}.csv".format(str(i).zfill(magnitude)), "w") as f:
            writer = csv.writer(f)

            headers = ['header_{}'.format(i) for i in range(1,15)]
            writer.writerow(headers)

            values = [[random.choice(list_of_random_values) for j in range(14)] for i in range(random.randint(MIN_ROWS_PER_FILE,MAX_ROWS_PER_FILE))]
            writer.writerows(values)

if __name__ == '__main__':
    generate_samples(25)