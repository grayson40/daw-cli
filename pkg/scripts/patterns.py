# Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
import json
from notes import *

# Creates pattern dictionary
Pattern = lambda name, channel, notes: {
    "name": name,
    "channel": channel,
    "notes": notes,
}

# Writes project patterns to json
def getPatterns(project):
    # Grab patterns
    project_patterns = project.patterns

    # Populate pattern list
    patterns_list = []
    for pattern in project_patterns:
        # Get notes and channel of pattern
        notes_list, rack_channel = getNotes(pattern)

        # Append pattern
        pattern_dict = Pattern(
            name=pattern.name, channel=rack_channel, notes=notes_list
        )
        patterns_list.append(pattern_dict)

    # Serializing json
    patterns_json = json.dumps(patterns_list, indent=4)

    # Write to json
    with open("./project/patterns.json", "w") as outfile:
        outfile.write(patterns_json)
