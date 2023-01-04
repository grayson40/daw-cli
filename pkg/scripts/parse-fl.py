# Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
import pyflp as fl
import os
import argparse
from channels import *
from mixer import *
from patterns import *
# from plugins import *

def main():
    # Parser obj
    parser = argparse.ArgumentParser()
    parser.add_argument('--input', type=str)
    parser.parse_args()
    print(parser.input)

    # # Filename 
    # file = "like-what.flp"

    # # Open and parse fl project
    # project = fl.parse(file)

    # # Make dir for json data
    # if not os.path.exists("./project"):
    #     os.mkdir("project")

    # # Get channels
    # getChannels(project)

    # # Get patterns
    # getPatterns(project)

    # # Get plugins
    # getPlugins(project)
    
    # # Get mixer
    # getMixer(project)


if __name__ == "__main__":
    main()
