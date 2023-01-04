# Copyright © 2022 Grayson Crozier <grayson40@gmail.com>

# Creates note dictionary
Note = lambda key, position, length: {
    "key": key,
    "position": position,
    "length": length,
}

# Returns notes and channel of pattern
def getNotes(pattern):
    # Note iterator
    notes_iter = iter(pattern)

    # Create list of notes
    notes_list = []
    rack_channel = 0
    for note in notes_iter:
        if rack_channel != note.rack_channel:
            rack_channel = note.rack_channel
        note_dict = Note(key=note.key, position=note.position, length=note.length)
        notes_list.append(note_dict)

    return notes_list, rack_channel