# Copyright © 2022 Grayson Crozier <grayson40@gmail.com>

import argparse
import pyflp as fl
import os


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--input', required=True)
    args = parser.parse_args()

    # Open staged write file
    staged = open("staged.txt", "w")

    if args.input == ".":
        files = []
        # Iterate over all project files in dir
        for file in os.listdir('./'):
            if file.split('.')[1] != 'flp':
                continue
            # Append filename
            files.append(file)


        for file in files:
            # Write filename to staged file
            staged.write(f'{file}\n')

            # # Parse fl project
            # project = fl.parse(file)
            # print(f'\nProject file: {file}\n\nChannels:')

            # # Grab channels
            # channels = project.channels
            # for channel in channels:
            #     print(channel.name)

        # Close staged write file
        staged.close()

    else:
        # Get input file
        input_file = args.input

        # Write filename to staged file
        staged.write(f'{input_file}\n')

        # # Parse fl project
        # project = fl.parse(input_file)
        # print(f'Project file: {input_file}\n\nChannels:')

        # # Grab channels
        # channels = project.channels
        # for channel in channels:
        #     print(channel.name)


if __name__ == "__main__":
    main()
